<template>
  <div class="common-manage">
    <el-card>
      <template #header>
        <div class="page-header">
          <span class="page-title">{{ pageTitle }}</span>
          <el-button type="primary" @click="handleAdd" v-if="showAdd">
            <el-icon><Plus /></el-icon>
            新增{{ itemName }}
          </el-button>
        </div>
      </template>

      <el-form :inline="true" :model="searchForm" class="search-form">
        <el-form-item label="关键词">
          <el-input v-model="searchForm.keyword" placeholder="请输入关键词" clearable @keyup.enter="handleSearch" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">
            <el-icon><Search /></el-icon>
            搜索
          </el-button>
          <el-button @click="handleReset">
            <el-icon><Refresh /></el-icon>
            重置
          </el-button>
        </el-form-item>
      </el-form>

      <el-table :data="tableData" v-loading="loading" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="名称" min-width="150" v-if="showName" />
        <el-table-column prop="title" label="标题" min-width="150" v-if="showTitle" />
        <el-table-column prop="hotel_name" label="酒店" min-width="120" v-if="showHotel" />
        <el-table-column prop="contact_name" label="联系人" width="100" v-if="showContact" />
        <el-table-column prop="total_price" label="金额" width="100" v-if="showPrice">
          <template #default="{ row }">
            ¥{{ row.total_price }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">
              {{ getStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="created_at" label="创建时间" width="180" />
        <el-table-column label="操作" fixed="right" width="200">
          <template #default="{ row }">
            <el-button type="primary" link @click="handleEdit(row)">
              <el-icon><View /></el-icon>
              查看
            </el-button>
            <el-button type="danger" link @click="handleDelete(row)" v-if="showDelete">
              <el-icon><Delete /></el-icon>
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        v-model:current-page="pagination.page"
        v-model:page-size="pagination.pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="pagination.total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="fetchData"
        @current-change="fetchData"
      />
    </el-card>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="600px">
      <el-form ref="formRef" :model="formData" :rules="formRules" label-width="100px">
        <el-form-item label="标题" prop="title" v-if="showTitle">
          <el-input v-model="formData.title" placeholder="请输入标题" />
        </el-form-item>
        <el-form-item label="描述" prop="description" v-if="showDesc">
          <el-input
            v-model="formData.description"
            type="textarea"
            :rows="4"
            placeholder="请输入描述"
          />
        </el-form-item>
        <el-form-item label="状态" prop="status" v-if="showStatusForm">
          <el-radio-group v-model="formData.status">
            <el-radio :label="1">正常/上架</el-radio>
            <el-radio :label="0">禁用/下架</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="submitLoading">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, getCurrentInstance } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import request from '@/utils/request'

const instance = getCurrentInstance()
const routePath = instance?.proxy?.$route?.path || ''

const pageConfig = {
  '/itinerary': { 
    title: '行程管理', 
    name: '行程', 
    api: '/admin/itineraries', 
    showName: false, 
    showTitle: true, 
    showDesc: false,
    showAdd: false,
    showDelete: true
  },
  '/daily-recommend': { 
    title: '每日推荐管理', 
    name: '推荐', 
    api: '/admin/daily-recommendations', 
    showName: false, 
    showTitle: true, 
    showDesc: true,
    showAdd: true,
    showDelete: true
  },
  '/hotel-stat': { 
    title: '酒店统计管理', 
    name: '统计', 
    api: '/admin/hotel-stats', 
    showName: false, 
    showHotel: true, 
    showPrice: true, 
    showStatusForm: false,
    showAdd: false,
    showDelete: false
  }
}

const currentConfig = pageConfig[routePath] || pageConfig['/itinerary']

const loading = ref(false)
const submitLoading = ref(false)
const dialogVisible = ref(false)
const isEdit = ref(false)
const formRef = ref(null)

const tableData = ref([])
const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

const searchForm = reactive({
  keyword: ''
})

const formData = reactive({
  id: 0,
  title: '',
  description: '',
  status: 1
})

const formRules = {
  title: [{ required: currentConfig.showTitle, message: '请输入标题', trigger: 'blur' }]
}

const pageTitle = computed(() => currentConfig.title)
const itemName = computed(() => currentConfig.name)
const showName = computed(() => currentConfig.showName)
const showTitle = computed(() => currentConfig.showTitle)
const showHotel = computed(() => currentConfig.showHotel)
const showContact = computed(() => currentConfig.showContact)
const showPrice = computed(() => currentConfig.showPrice)
const showDesc = computed(() => currentConfig.showDesc)
const showStatusForm = computed(() => currentConfig.showStatusForm !== false)
const showAdd = computed(() => currentConfig.showAdd !== false)
const showDelete = computed(() => currentConfig.showDelete !== false)

const dialogTitle = computed(() => isEdit.value ? '编辑' + currentConfig.name : '新增' + currentConfig.name)

const getStatusType = (status) => {
  const types = { 0: 'danger', 1: 'success', 2: 'warning', 3: 'info' }
  return types[status] || 'info'
}

const getStatusText = (status) => {
  const texts = { 0: '禁用/下架', 1: '正常/上架', 2: '进行中', 3: '已完成' }
  return texts[status] || '未知'
}

const fetchData = async () => {
  loading.value = true
  try {
    const params = {
      page: pagination.page,
      page_size: pagination.pageSize,
      keyword: searchForm.keyword
    }
    const res = await request.get(currentConfig.api, { params })
    tableData.value = res.data?.items || []
    pagination.total = res.data?.total || 0
  } catch (error) {
    console.error('Failed to fetch data:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.page = 1
  fetchData()
}

const handleReset = () => {
  searchForm.keyword = ''
  handleSearch()
}

const handleAdd = () => {
  isEdit.value = false
  Object.assign(formData, {
    id: 0,
    title: '',
    description: '',
    status: 1
  })
  dialogVisible.value = true
}

const handleEdit = (row) => {
  isEdit.value = true
  Object.assign(formData, {
    id: row.id,
    title: row.title || '',
    description: row.description || '',
    status: row.status
  })
  dialogVisible.value = true
}

const handleDelete = (row) => {
  ElMessageBox.confirm(`确定要删除该${currentConfig.name}吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await request.delete(`${currentConfig.api}/${row.id}`)
      ElMessage.success('删除成功')
      fetchData()
    } catch (error) {
      console.error('Failed to delete:', error)
    }
  }).catch(() => {})
}

const handleSubmit = async () => {
  submitLoading.value = true
  try {
    if (isEdit.value) {
      await request.put(`${currentConfig.api}/${formData.id}`, formData)
      ElMessage.success('更新成功')
    } else {
      await request.post(currentConfig.api, formData)
      ElMessage.success('创建成功')
    }
    dialogVisible.value = false
    fetchData()
  } catch (error) {
    console.error('Failed to submit:', error)
  } finally {
    submitLoading.value = false
  }
}

onMounted(() => {
  fetchData()
})
</script>

<style scoped>
.common-manage {
  padding: 0;
}
</style>
