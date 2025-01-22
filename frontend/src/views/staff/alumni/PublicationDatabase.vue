<template>
  <div class="statistics-container">
    <!-- Top Section: Filters -->
    <div class="filter-section">
      <el-row :gutter="20">
        <!-- Date Range Picker -->
        <el-col :span="8">
          <el-date-picker
              v-model="dateRange"
              type="daterange"
              unlink-panels
              range-separator="To"
              start-placeholder="Start date"
              end-placeholder="End date"
              format="DD-MM-YYYY"
              :shortcuts="shortcuts"
              @change="handleDateChange"
          />
        </el-col>

        <!-- Search Box -->
        <el-col :span="8">
          <el-input
              v-model="searchQuery"
              placeholder="Search publications"
              @input="handleSearch"
          >
            <template #prefix>
              <el-icon>
                <Search/>
              </el-icon>
            </template>
          </el-input>
        </el-col>
      </el-row>

      <el-row :gutter="20" class="mt-3">
        <!-- Publication Type Filter -->
        <el-col :span="8">
          <el-checkbox-group v-model="selectedTypes" @change="handleFiltersChange">
            <el-checkbox label="ISI WOS">ISI WOS</el-checkbox>
            <el-checkbox label="Scopus">Scopus</el-checkbox>
            <el-checkbox label="Book">Book</el-checkbox>
            <el-checkbox label="Chapter of Book">Chapter of Book</el-checkbox>
            <el-checkbox label="Journal A">Journal A</el-checkbox>
            <el-checkbox label="Journal B">Journal B</el-checkbox>
            <el-checkbox label="Others">Others</el-checkbox>
          </el-checkbox-group>
        </el-col>

        <!-- Status Filter -->
        <el-col :span="8">
          <el-checkbox-group v-model="selectedStatus" @change="handleFiltersChange">
            <el-checkbox label="Accepted">Accepted</el-checkbox>
            <el-checkbox label="Published">Published</el-checkbox>
            <el-checkbox label="Draft">Draft</el-checkbox>
            <el-checkbox label="Under Review">Under Review</el-checkbox>
          </el-checkbox-group>
        </el-col>

        <!-- Quartile Filter -->
        <el-col :span="8">
          <el-checkbox-group v-model="selectedQuartiles" @change="handleFiltersChange">
            <el-checkbox label="Q1">Q1</el-checkbox>
            <el-checkbox label="Q2">Q2</el-checkbox>
            <el-checkbox label="Q3">Q3</el-checkbox>
            <el-checkbox label="Q4">Q4</el-checkbox>
          </el-checkbox-group>
        </el-col>
      </el-row>
    </div>

    <!-- Middle Section: ECharts Pie Charts -->
    <div class="chart-section">
      <el-row :gutter="20">
        <el-col :span="8">
          <div ref="typeChartRef" style="height: 400px;"></div>
        </el-col>
        <el-col :span="8">
          <div ref="statusChartRef" style="height: 400px;"></div>
        </el-col>
        <el-col :span="8">
          <div ref="quartileChartRef" style="height: 400px;"></div>
        </el-col>
      </el-row>
    </div>

    <!-- Bottom Section: Publication Table -->
    <div class="table-section">
      <el-table
          :data="tableData"
          style="width: 100%"
          border
          v-loading="loading"
      >
        <el-table-column
            prop="alumni_name"
            label="Alumni Name"
            min-width="200"
        />
        <el-table-column
            prop="article_title"
            label="Article Title"
            min-width="800"
            show-overflow-tooltip
        />
        <el-table-column
            label="Journal Title"
            min-width="150"
            show-overflow-tooltip>
          <template #default="scope">
            {{ scope.row.journal_title || '-' }}
          </template>
        </el-table-column>
        <el-table-column
            prop="quartile"
            label="Quartile"
            width="100"
            align="center"
        />
        <el-table-column
            label="Authors"
            min-width="180"
            show-overflow-tooltip>
          <template #default="scope">
            {{ scope.row.authors || '-' }}
          </template>
        </el-table-column>
        <el-table-column
            prop="publication_type"
            label="Publication Type"
            width="150"
            align="center"
        />
        <el-table-column
            prop="status"
            label="Status"
            width="120"
            align="center"
        >
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">
              {{ row.status }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column
            prop="accepted_date"
            label="Accepted Date"
            width="120"
            align="center"
        >
          <template #default="{ row }">
            {{ formatDateFromDB(row.accepted_date) || '-' }}
          </template>
        </el-table-column>
        <el-table-column
            label="Corresponding Authors"
            min-width="180"
            show-overflow-tooltip>
          <template #default="scope">
            {{ scope.row.corresponding_authors || '-' }}
          </template>
        </el-table-column>

      </el-table>

      <div class="pagination-container">
        <el-pagination
            v-model:current-page="currentPage"
            v-model:page-size="pageSize"
            :page-sizes="[10, 20, 50, 100]"
            :total="total"
            layout="total, sizes, prev, pager, next, jumper"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import {ref, watch, onMounted, onUnmounted} from 'vue'
import {use} from 'echarts/core'
import {CanvasRenderer} from 'echarts/renderers'
import {PieChart} from 'echarts/charts'
import {LegendComponent, TooltipComponent} from 'echarts/components'
import * as echarts from 'echarts'
import {apiGetStatistics} from '@/apis/alumni/profile/publication.js'

// Initialize ECharts
use([CanvasRenderer, PieChart, LegendComponent, TooltipComponent])

// Chart references
const typeChartRef = ref(null)
const statusChartRef = ref(null)
const quartileChartRef = ref(null)
let typeChart = null
let statusChart = null
let quartileChart = null

const statusMap = {
  'Published': 'success',
  'Accepted': 'primary',
  'Draft': 'info',
  'Under Review': 'warning'
}

const loading = ref(false)
const tableData = ref([])
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)

// Filter states
const dateRange = ref([])
const searchQuery = ref('')
const selectedTypes = ref([])
const selectedStatus = ref([])
const selectedQuartiles = ref([])

const getStatusType = (status) => {
  return statusMap[status] || ''
}


// Date shortcuts
const shortcuts = [
  {
    text: 'Last month',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
      return [start, end]
    },
  },
  {
    text: 'Last 3 months',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
      return [start, end]
    },
  },
  {
    text: 'Last year',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 365)
      return [start, end]
    },
  },
]

const formatDateFromDB = (dbDate) => {
  if (!dbDate ||
      dbDate === '0001-01-01T00:00:00Z' ||
      dbDate === '1970-01-01T00:00:00Z') {
    return '-';
  }

  try {
    const date = dbDate.split('T')[0];
    const [year, month, day] = date.split('-');

    if (!year || !month || !day) {
      return '-';
    }

    return `${day}-${month}-${year}`;
  } catch (error) {
    return '-';
  }
}

// to db format (2024-08-30)
const formatDateForDB = (date) => {
  if (!date) return '';

  try {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');

    return `${year}-${month}-${day}`;
  } catch (error) {
    console.error('Date formatting error:', error);
    return '';
  }
}

// Initialize charts
const initCharts = () => {
  if (typeChartRef.value) {
    typeChart = echarts.init(typeChartRef.value)
  }
  if (statusChartRef.value) {
    statusChart = echarts.init(statusChartRef.value)
  }
  if (quartileChartRef.value) {
    quartileChart = echarts.init(quartileChartRef.value)
  }
  updateCharts()
}

// Create chart option
const createChartOption = (title, data) => {
  return {
    title: {
      text: title,
      left: 'center'
    },
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b}: {c} ({d}%)'
    },
    legend: {
      orient: 'vertical',
      left: 'left',
      top: 'middle'
    },
    series: [
      {
        name: title,
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 2
        },
        label: {
          show: true,
          formatter: '{b}: {c} ({d}%)'
        },
        emphasis: {
          label: {
            show: true,
            fontSize: '16',
            fontWeight: 'bold'
          }
        },
        data: Object.entries(data).map(([name, value]) => ({
          name,
          value
        }))
      }
    ]
  }
}

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toISOString().split('T')[0]
}


// Update charts with new data
const updateCharts = async () => {
  loading.value = true
  const params = getQueryParams()
  try {
    const response = await apiGetStatistics(params)
    if (response.data.code === 200 && response.data.meta.statistics) {
      const stats = response.data.meta.statistics

      // Update publication types chart
      typeChart?.setOption(createChartOption(
          'Publication Types',
          stats.publicationTypeCount
      ))

      // Update status chart
      statusChart?.setOption(createChartOption(
          'Publication Status',
          stats.statusCount
      ))

      // Update quartile chart
      quartileChart?.setOption(createChartOption(
          'Publication Quartiles',
          stats.quartileCount
      ))

      // Update table data
      tableData.value = response.data.data
      total.value = response.data.meta.pagination.total
    }
  } catch (error) {
    console.error('Failed to fetch statistics:', error)
  } finally {
    loading.value = false
  }
}

// Get query parameters
const getQueryParams = () => {
  const params = {
    page_Size: pageSize.value,
    current_page: currentPage.value
  }

  if (selectedStatus.value.length > 0) {
    params.status = selectedStatus.value
  }

  if (selectedTypes.value.length > 0) {
    params.publication_type = selectedTypes.value
  }

  if (selectedQuartiles.value.length > 0) {
    params.quartile = selectedQuartiles.value
  }

  if (dateRange.value?.length === 2) {
    const [startDate, endDate] = dateRange.value
    params.accepted_date_start = formatDateForDB(startDate)
    params.accepted_date_end = formatDateForDB(endDate)
  }

  if (searchQuery.value) {
    params.search = searchQuery.value
  }

  return params
}


// Event handlers
const handleDateChange = (dates) => {
  if (dates) {
    dateRange.value = dates
  } else {
    dateRange.value = []
  }
  currentPage.value = 1
  updateCharts()
}

const handleSearch = () => {
  updateCharts()
}

const handleFiltersChange = () => {
  updateCharts()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  currentPage.value = 1
  updateCharts()
}

const handleCurrentChange = (val) => {
  currentPage.value = val
  updateCharts()
}

// Handle window resize
const handleResize = () => {
  typeChart?.resize()
  statusChart?.resize()
  quartileChart?.resize()
}

watch(
    dateRange,
    (newValue) => {
      if (newValue) {
        currentPage.value = 1
        updateCharts()
      }
    },
    {deep: true, immediate: true}
)

watch(
    [dateRange, searchQuery, selectedTypes, selectedStatus, selectedQuartiles],
    () => {
      currentPage.value = 1
      updateCharts()
    },
    {deep: true}
)

// Lifecycle hooks
onMounted(() => {
  initCharts()
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  typeChart?.dispose()
  statusChart?.dispose()
  quartileChart?.dispose()
  window.removeEventListener('resize', handleResize)
})

</script>

<style scoped>
.statistics-container {
  padding: 20px;
}

.filter-section {
  margin-bottom: 20px;
}

.chart-section {
  margin-bottom: 20px;
}

.table-section {
  margin-bottom: 20px;
}

.mt-3 {
  margin-top: 1rem;
}
</style>
