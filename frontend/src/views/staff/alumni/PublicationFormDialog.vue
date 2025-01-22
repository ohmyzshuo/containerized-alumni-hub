<template>
  <el-dialog
      v-model="visible"
      :title="isEdit ? 'Edit Publication' : 'Add Publication'"
      width="50%"
  >
    <el-form
        ref="formRef"
        :model="form"
        :rules="rules"
        label-width="140px"
        label-position="right"
    >
      <el-form-item label="Title" prop="article_title">
        <el-input v-model="form.article_title"/>
      </el-form-item>

      <el-form-item label="Journal Title" prop="journal_title">
        <el-input v-model="form.journal_title"/>
      </el-form-item>

      <el-form-item label="Publication Type" prop="publication_type">
        <el-select v-model="form.publication_type" placeholder="Select type">
          <el-option label="ISI WOS" value="ISI WOS"/>
          <el-option label="Scopus" value="Scopus"/>
          <el-option label="Chapter of Book" value="Chapter of Book"/>
          <el-option label="Journal A/B" value="Journal A/B"/>
          <el-option label="Conference" value="Conference"/>
          <el-option label="Book" value="Book"/>
          <el-option label="Other" value="Other"/>
        </el-select>
      </el-form-item>

      <el-form-item label="Authors" prop="authors">
        <el-input v-model="form.authors" type="textarea"/>
      </el-form-item>

      <el-form-item label="Corresponding Authors" prop="corresponding_authors">
        <el-input v-model="form.corresponding_authors" type="textarea"/>
      </el-form-item>

      <el-form-item label="Quartile" prop="quartile">
        <el-select v-model="form.quartile" placeholder="Select quartile">
          <el-option label="Q1" value="Q1"/>
          <el-option label="Q2" value="Q2"/>
          <el-option label="Q3" value="Q3"/>
          <el-option label="Q4" value="Q4"/>
        </el-select>
      </el-form-item>

      <el-form-item label="Status" prop="status">
        <el-select v-model="form.status" placeholder="Select status">
          <el-option label="Published" value="Published"/>
          <el-option label="Accepted" value="Accepted"/>
          <el-option label="Under Review" value="Under Review"/>
          <el-option label="Draft" value="Draft"/>
        </el-select>
      </el-form-item>

      <el-form-item
          label="Accepted Date"
          prop="accepted_date"
          v-if="showDatePicker"
      >
        <el-date-picker
            v-model="form.accepted_date"
            type="date"
            placeholder="Select date"
            format="YYYY-MM-DD"
        />
      </el-form-item>
    </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="visible = false">Cancel</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="loading">
          {{ isEdit ? 'Update' : 'Create' }}
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import {ref, computed, defineProps, defineEmits, watch} from 'vue'

const props = defineProps({
  modelValue: Boolean,
  publication: {
    type: Object,
    default: () => ({})
  },
  isEdit: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'submit'])

const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const formRef = ref(null)
const loading = ref(false)

const defaultForm = {
  article_title: '',
  journal_title: '',
  publication_type: '',
  authors: '',
  corresponding_authors: '',
  quartile: '',
  status: '',
  accepted_date: ''
}

const form = ref({...defaultForm})

const showDatePicker = computed(() => {
  return ['Published', 'Accepted'].includes(form.value.status)
})

watch(() => form.value.status, (newStatus) => {
  if (!['Published', 'Accepted'].includes(newStatus)) {
    form.value.accepted_date = ''
  }
})

const rules = {
  article_title: [{required: true, message: 'Please input title', trigger: 'blur'}],
  publication_type: [{required: true, message: 'Please select type', trigger: 'change'}],
  status: [{required: true, message: 'Please select status', trigger: 'change'}],
  accepted_date: [
    {
      required: true,
      message: 'Please select date',
      trigger: 'change',
      validator: (rule, value, callback) => {
        if (['Published', 'Accepted'].includes(form.value.status) && !value) {
          callback(new Error('Please select date'))
        } else {
          callback()
        }
      }
    }
  ]
}

watch(() => props.publication, (newVal) => {
  if (newVal && Object.keys(newVal).length > 0) {
    form.value = {...newVal}
    if (form.value.accepted_date) {
      form.value.accepted_date = new Date(form.value.accepted_date)
    }
  } else {
    form.value = {...defaultForm}
  }
}, {deep: true, immediate: true})

const handleSubmit = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true
      try {
        const submitData = {
          ...form.value,
          "accepted_date": form.value.accepted_date === '' ? '1970-01-01' : form.value.accepted_date,
        }
        if (submitData.accepted_date) {
          submitData.accepted_date = new Date(submitData.accepted_date).toISOString()
        }
        emit('submit', submitData)
        visible.value = false
      } finally {
        loading.value = false
      }
    }
  })
}
</script>
