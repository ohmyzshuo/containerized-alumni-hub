<template>
  <div class="content-table-container">
    <div class="header mb-6">
      <h1 class="text-2xl font-bold text-gray-800">Content Management</h1>
    </div>
    <div class="content-table">
      <!-- Search and Action Area -->
      <div class="mb-4 flex justify-between items-center content-header">
        <div class="left-section">
          <el-button
              type="primary"
              @click="handleSendNewsletter"
              :disabled="selectedContents.length === 0"
          >
            Send Newsletter
          </el-button>
          <el-button type="primary" @click="showCreateDialog">Create</el-button>
        </div>
        <div class="right-section">
          <el-input
              v-model="searchQuery"
              placeholder="Search contents..."
              @input="handleSearch"
          />
        </div>
      </div>

      <el-table
          :data="contents"
          style="width: 100%"
          v-loading="loading"
          class="main-table"
          @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55"/>
        <el-table-column label="Type" min-width="110">
          <template #default="{ row }">
            <el-tag :type="getContentTypeTag(row.content_type)">
              {{ getContentTypeLabel(row.content_type) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="Title" prop="title" min-width="150" show-overflow-tooltip/>
        <el-table-column label="Description" prop="description" min-width="200" show-overflow-tooltip/>
        <el-table-column label="Venue" prop="venue" min-width="100">
          <template #default="{ row }">
            {{ row.content_type !== 1 ? '-' : row.venue }}
          </template>
        </el-table-column>
        <el-table-column label="Visibility" min-width="80">
          <template #default="{ row }">
            <el-tag :type="row.is_hidden ? 'warning' : 'info'">
              {{ row.is_hidden ? 'Hidden' : 'Visible' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="Participant Quota" min-width="80">
          <template #default="{ row }">
            {{
              row.content_type !== 1 ? '-' : (row.participant_quota === 0 || row.participant_quota === undefined ? 'Unlimited' : row.participant_quota)
            }}
          </template>
        </el-table-column>
        <el-table-column label="Attachments" min-width="100">
          <template #default="{ row }">
            <span>{{ row.attachments.length }}</span>
          </template>
        </el-table-column>
        <el-table-column label="Created By" prop="created_by_name" min-width="100"/>
        <el-table-column label="Created Time" min-width="160">
          <template #default="{ row }">
            {{ new Date(row.created_at).toLocaleString() }}
          </template>
        </el-table-column>
        <el-table-column label="Operation" width="300">
          <template #default="scope">
            <el-button
                size="small"
                type="success"
                @click="handleShowDetails(scope.row)"
                :disabled="scope.row.content_type !== 1"
            >
              Participation
            </el-button>
            <el-button size="small" @click="handleEdit(scope.row)">Edit</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.row)">Delete</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-container">
        <el-pagination
            v-model:current-page="currentPage"
            v-model:page-size="pageSize"
            :total="total"
            :page-sizes="[10, 15, 20, 30]"
            layout="total, sizes, prev, pager, next"
            @size-change="handleSizeChange"
            @current-change="handlePageChange"
        />
      </div>

      <el-dialog
          v-model="dialogVisible"
          :title="isEdit ? 'Edit Content' : 'Create Content'"
          width="50%"
          @closed="handleDialogClose"
      >
        <el-form :model="formData" label-width="120px">
          <el-form-item label="Content Type">
            <el-select v-model="formData.content_type">
              <el-option
                  v-for="(label, value) in contentTypes"
                  :key="value"
                  :label="label"
                  :value="Number(value)"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="Title">
            <el-input v-model="formData.title"/>
          </el-form-item>
          <el-form-item label="Description">
            <el-input v-model="formData.description" type="textarea" rows="4"/>
          </el-form-item>
          <template v-if="formData.content_type === 1">
            <el-form-item label="Venue">
              <el-input v-model="formData.venue"/>
            </el-form-item>
            <el-form-item label="Participant Quota">
              <el-input-number
                  v-model="formData.participant_quota"
                  :min="0"
                  :controls="true"
                  placeholder="0 means unlimited"
              />

            </el-form-item>
          </template>
          <el-form-item label="Attachments">
            <el-upload
                :http-request="customUpload"
                :on-remove="handleRemove"
                :on-preview="handlePreview"
                :before-upload="beforeUpload"
                multiple
                :file-list="fileList"
                :limit="5"
            >
              <el-button type="primary">Upload Files</el-button>
              <template #tip>
                <div class="el-upload__tip">
                  Supported file types: IMAGES, PDF, DOC, DOCX, XLS, XLSX. Max size: 10MB
                </div>
              </template>
            </el-upload>
          </el-form-item>
          <el-form-item label="Visibility">
            <el-switch
                v-model="formData.is_hidden"
                active-text="Hidden"
                inactive-text="Visible"
            />
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">Cancel</el-button>
            <el-button type="primary" @click="handleSubmit">
              Confirm
            </el-button>
          </span>
        </template>
      </el-dialog>

      <el-dialog
          v-model="previewDialogVisible"
          title="Preview"
          @close="handlePreviewClose"
      >
        <template v-if="previewUrl">
          <img
              v-if="previewUrl.match(/\.(jpg|jpeg|png|gif)$/i)"
              :src="previewUrl"
              style="max-width: 100%; max-height: 70vh;"
          />
          <iframe
              v-else-if="previewUrl.match(/\.pdf$/i)"
              :src="previewUrl"
              style="width: 100%; height: 70vh;"
          ></iframe>
        </template>
      </el-dialog>

      <el-dialog
          v-model="showDetailsDialog"
          title="Participation Details"
          width="700px"
          destroy-on-close
      >
        <el-table
            v-loading="detailsLoading"
            :data="participationDetails"
            stripe
            style="width: 100%"
        >
          <el-table-column prop="name" label="Name"/>
          <el-table-column prop="matric_no" label="Student ID"/>
          <el-table-column prop="email" label="Email"/>
          <el-table-column prop="participant_status" label="Status">
            <template #default="scope">
              <el-tag :type="getStatusType(scope.row.participant_status)">
                {{ getStatusText(scope.row.participant_status) }}
              </el-tag>
            </template>
          </el-table-column>
          <el-table-column label="Operation" width="120">
            <template #default="scope">
              <el-button
                  size="small"
                  type="primary"
                  @click="handleChangeStatus(scope.row)"
              >
                Change
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-dialog>

      <el-dialog
          v-model="showStatusDialog"
          title="Change Status"
          width="400px"
          append-to-body
      >
        <el-form :model="statusForm" label-width="100px">
          <el-form-item label="New Status">
            <el-select v-model="statusForm.to_status" placeholder="Select status">
              <el-option
                  v-for="(value, key) in CHANGEABLE_STATUS_MAP"
                  :key="key"
                  :label="value.text"
                  :value="Number(key)"
              />
            </el-select>
          </el-form-item>
        </el-form>
        <template #footer>
      <span class="dialog-footer">
        <el-button @click="showStatusDialog = false">Cancel</el-button>
        <el-button type="primary" @click="submitStatusChange" :loading="statusChangeLoading">
          Confirm
        </el-button>
      </span>
        </template>
      </el-dialog>

      <!-- Newsletter Dialog -->
      <el-dialog
          v-model="newsletterDialogVisible"
          title="Send Newsletter"
          width="50%"
      >
        <div class="newsletter-search">
          <el-input
              v-model="alumniSearch"
              placeholder="Search alumni..."
              @input="handleAlumniSearch"
          />
        </div>

        <el-table
            :data="alumniList"
            @selection-change="handleAlumniSelectionChange"
        >
          <el-table-column type="selection" width="55"/>
          <el-table-column prop="name" label="Name"/>
          <el-table-column prop="matric_no" label="Matric No"/>
          <el-table-column prop="email" label="Email"/>
        </el-table>

        <template #footer>
        <span class="dialog-footer">
          <el-button @click="newsletterDialogVisible = false">Cancel</el-button>
          <el-button type="primary" @click="confirmSendNewsletter">
            Send
          </el-button>
        </span>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script setup>
import {ref, onMounted} from 'vue'
import {ElMessage, ElMessageBox} from 'element-plus'
import {
  apiGetContents,
  apiEditContent,
  apiCreateContent,
  apiDeleteContent,
  apiGetEventParticipation,
  apiChangeParticipantStatus,
  apiGetContentDetails,
  apiSendNewsletter,
} from '@/apis/staff/content/content.js'
import {apiGetAlumni} from "@/apis/alumni/alumni.js";

const contentTypes = {
  1: "Event",
  2: "Announcement",
  3: "Advertisement",
  0: "Other"
}

const STATUS_MAP = {
  0: {text: 'No Response', type: 'info'},
  1: {text: 'Not Interested', type: 'danger'},
  2: {text: 'Interested but Unavailable', type: 'warning'},
  3: {text: 'Registered', type: 'primary'},
  4: {text: 'Attended', type: 'success'},
  5: {text: 'Absent', type: 'danger'}
}

const CHANGEABLE_STATUS_MAP = {
  2: {text: 'Interested but Unavailable', type: 'warning'},
  3: {text: 'Registered', type: 'primary'},
  4: {text: 'Attended', type: 'success'},
  5: {text: 'Absent', type: 'danger'}
}

const contents = ref([])
const loading = ref(false)
const searchQuery = ref('')
const currentPage = ref(1)
const pageSize = ref(12)
const total = ref(0)

const selectedContents = ref([])
const newsletterDialogVisible = ref(false)
const alumniSearch = ref('')
const alumniList = ref([])
const selectedAlumni = ref([])

const showDetailsDialog = ref(false)
const participationDetails = ref([])
const detailsLoading = ref(false)

const dialogVisible = ref(false)
const isEdit = ref(false)
const formData = ref({
  title: '',
  description: '',
  venue: '',
  is_hidden: false,
  created_by_name: '',
  participant_quota: 0,
  content_type: 0,
  attachments: [],
})

const showStatusDialog = ref(false)
const statusChangeLoading = ref(false)
const statusForm = ref({
  alumni_id: null,
  content_id: null,
  to_status: null
})
const currentContentId = ref(null)

const fileList = ref([])
const previewDialogVisible = ref(false)
const previewUrl = ref('')

const getStatusText = (status) => {
  return STATUS_MAP[status]?.text || 'Unknown'
}

const getStatusType = (status) => {
  return STATUS_MAP[status]?.type || 'info'
}

function getContentTypeLabel(contentType) {
  return contentTypes[contentType] || "Other"
}

function getContentTypeTag(contentType) {
  const tagTypes = {
    1: 'success',
    2: 'warning',
    3: 'danger',
    0: 'info'
  }
  return tagTypes[contentType] || 'info'
}

const getContentsList = async () => {
  loading.value = true
  try {
    const params = {
      page: currentPage.value,
      page_size: pageSize.value,
      search: searchQuery.value || undefined
    }
    const res = await apiGetContents(params)
    if (res.data.code === 200) {
      contents.value = res.data.data
      total.value = res.data.meta.total
    }
  } catch (error) {
    ElMessage.error('Failed to fetch contents')
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  currentPage.value = 1
  getContentsList()
}

const handlePageChange = (page) => {
  currentPage.value = page
  getContentsList()
}

const handleSizeChange = (size) => {
  pageSize.value = size
  currentPage.value = 1
  getContentsList()
}

const showCreateDialog = () => {
  isEdit.value = false
  formData.value = {
    title: '',
    description: '',
    venue: '',
    is_hidden: false,
    created_by_name: '',
    participant_quota: 0,
    content_type: 0,
    attachments: [],
  }
  dialogVisible.value = true
}


const handleEdit = async (row) => {
  try {
    const response = await apiGetContentDetails(row.id);

    const formattedAttachments = response.data.data.attachments.map(attachment => ({
      name: attachment.original_name,
      url: `${import.meta.env.VITE_API_BASE_URL}/${attachment.attachment_path}`,
      response: attachment,
      status: 'success'
    }));

    fileList.value = formattedAttachments;

    formData.value = {
      ...response.data.data.content,
      attachments: formattedAttachments
    };

    dialogVisible.value = true;
    isEdit.value = true
  } catch (error) {
    console.error('Error getting content detail:', error);
    ElMessage.error('Failed to get content detail');
  }
};

const handleDelete = async (row) => {
  try {
    await ElMessageBox.confirm('Are you sure you want to delete this content?', 'Warning', {
      type: 'warning'
    })
    await apiDeleteContent(row.id)
    ElMessage.success('Successfully deleted')
    await getContentsList()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('Failed to delete')
    }
  }
}

const handleShowDetails = async (row) => {
  if (row.content_type !== 1) return

  currentContentId.value = row.id
  showDetailsDialog.value = true
  detailsLoading.value = true

  try {
    const res = await apiGetEventParticipation({
      content_id: row.id
    })
    if (res.data.code === 200) {
      participationDetails.value = res.data.data
    }
  } catch (error) {
    console.error('Error fetching details:', error)
    ElMessage.error('Failed to fetch participation details')
  } finally {
    detailsLoading.value = false
  }
}

const handleChangeStatus = (row) => {
  statusForm.value = {
    alumni_id: row.id,
    content_id: currentContentId.value,
    to_status: row.participant_status
  }
  showStatusDialog.value = true
}


const submitStatusChange = async () => {
  statusChangeLoading.value = true
  try {
    const res = await apiChangeParticipantStatus(statusForm.value)
    if (res.data.code === 200) {
      ElMessage.success('Status updated successfully')
      showStatusDialog.value = false
      await handleShowDetails({id: currentContentId.value, content_type: 1})
    }
  } catch (error) {
    console.error('Error changing status:', error)
    ElMessage.error('Failed to update status')
  } finally {
    statusChangeLoading.value = false
  }
}

const customUpload = async (options) => {
  try {
    const file = options.file;
    formData.value.attachments.push(file);
    fileList.value.push({
      name: file.name,
      size: file.size,
      raw: file
    });
    options.onSuccess();
  } catch (error) {
    options.onError(error);
    ElMessage.error('Upload failed');
  }
};

const handleRemove = (file) => {
  const index = formData.value.attachments.findIndex(
      (item) => item.name === file.name && item.size === file.size
  );
  if (index > -1) {
    formData.value.attachments.splice(index, 1);
  }
  const fileIndex = fileList.value.findIndex(
      (item) => item.name === file.name && item.size === file.size
  );
  if (fileIndex > -1) {
    fileList.value.splice(fileIndex, 1);
  }
};

const handlePreview = (file) => {
  const fileType = file.name.split('.').pop().toLowerCase();
  const fileUrl = file.raw ? URL.createObjectURL(file.raw) : file.url;

  if (['jpg', 'jpeg', 'png', 'gif'].includes(fileType)) {
    previewUrl.value = fileUrl;
    previewDialogVisible.value = true;
  } else if (fileType === 'pdf') {
    previewUrl.value = fileUrl;
    previewDialogVisible.value = true;
  } else {
    if (file.raw) {
      const link = document.createElement('a');
      link.href = URL.createObjectURL(file.raw);
      link.download = file.name;
      link.click();
    } else {
      window.open(fileUrl, '_blank');
    }
  }
};

const handlePreviewClose = () => {
  if (previewUrl.value && previewUrl.value.startsWith('blob:')) {
    URL.revokeObjectURL(previewUrl.value);
  }
  previewUrl.value = '';
  previewDialogVisible.value = false;
};
const beforeUpload = (file) => {
  const allowedTypes = [
    'image/jpeg',
    'image/png',
    'image/gif',
    'application/pdf',
    'application/msword',
    'application/vnd.openxmlformats-officedocument.wordprocessingml.document',
    'application/vnd.ms-excel',
    'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
  ];

  const isValidType = allowedTypes.includes(file.type);
  if (!isValidType) {
    ElMessage.error('Unsupported file type!');
    return false;
  }

  const maxSize = 10 * 1024 * 1024;
  if (file.size > maxSize) {
    ElMessage.error('File size cannot exceed 10MB!');
    return false;
  }

  return true;
};

const handleSubmit = async () => {
  try {
    const submitData = new FormData();

    for (const key in formData.value) {
      if (key === 'attachments') continue;

      if (typeof formData.value[key] === 'boolean') {
        submitData.append(key, formData.value[key] ? '1' : '0');
      } else if (typeof formData.value[key] === 'number') {
        submitData.append(key, formData.value[key].toString());
      } else if (formData.value[key] != null) {
        submitData.append(key, formData.value[key]);
      }
    }

    if (formData.value.attachments && formData.value.attachments.length > 0) {
      formData.value.attachments.forEach((file, index) => {
        submitData.append('attachments', file);
      });
    }

    if (isEdit.value) {
      await apiEditContent(submitData, formData.value.id);
      ElMessage.success('Successfully updated');
    } else {
      await apiCreateContent(submitData);
      ElMessage.success('Successfully created');
    }
    dialogVisible.value = false;
    await getContentsList();

    fileList.value = [];
    formData.value.attachments = [];
  } catch (error) {
    console.error('Submit error:', error);
    ElMessage.error(isEdit.value ? 'Failed to update' : 'Failed to create');
  }
};
const handleDialogClose = () => {
  fileList.value = [];
  formData.value.attachments = [];
};

// Content selection handler
const handleSelectionChange = (selection) => {
  selectedContents.value = selection.map(item => item.id)
}

// Open newsletter dialog
const handleSendNewsletter = () => {
  newsletterDialogVisible.value = true
  alumniList.value = []
  selectedAlumni.value = []
  alumniSearch.value = ''
}

// Search alumni
const handleAlumniSearch = async () => {
  if (alumniSearch.value) {
    try {
      const res = await apiGetAlumni(alumniSearch.value)
      if (res.data.code === 200) {
        alumniList.value = res.data.data
      }
    } catch (error) {
      ElMessage.error('Failed to search alumni')
    }
  }
}

// Alumni selection handler
const handleAlumniSelectionChange = (selection) => {
  selectedAlumni.value = selection.map(item => item.id)
}

// Confirm and send newsletter
const confirmSendNewsletter = async () => {
  if (selectedAlumni.value.length === 0) {
    ElMessage.warning('Please select at least one recipient')
    return
  }

  try {
    const data = {
      alumni_ids: selectedAlumni.value,
      content_ids: selectedContents.value
    }
    const res = await apiSendNewsletter(data)
    if (res.data.code === 200) {
      ElMessage.success('Newsletter sent successfully')
      newsletterDialogVisible.value = false
    }
  } catch (error) {
    ElMessage.error('Failed to send newsletter')
  }
}


onMounted(() => {
  getContentsList()
})
</script>

<style scoped>
.content-table-container {
  width: 100%;
  height: 100vh;
  padding: 20px;
  box-sizing: border-box;
  background-color: #f5f7fa;
}

.content-table {
  height: 100%;
  background-color: #fff;
  border-radius: 4px;
  padding: 20px;
  display: flex;
  flex-direction: column;
}

.content-header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}

.left-section {
  display: flex;
  gap: 10px;
}

.newsletter-search {
  margin-bottom: 20px;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

.main-table {
  flex: 1;
  margin-bottom: 20px;
}

.pagination-container {
  display: flex;
  justify-content: flex-end;
  padding: 10px 0;
}

:deep(el-table) {
  height: 100%;
}

:deep(.el-table__inner-wrapper) {
  height: 100%;
}

:deep(.el-table__body-wrapper) {
  overflow-y: auto;
}
</style>
