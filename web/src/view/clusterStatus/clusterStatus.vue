<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="created_at">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="cluster_name">
         <el-input v-model="searchInfo.cluster_name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="pmu_name">
         <el-input v-model="searchInfo.pmu_name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="working_status">
         <el-select v-model="searchInfo.working_status" clearable placeholder="请选择">
            <el-option
              key="available"
              label="available"
              value="available">
            </el-option>
            <el-option
              key="booked"
              label="booked"
              value="booked">
            </el-option>
            <el-option
              key="working"
              label="working"
              value="working">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="health_status">
         <el-input v-model="searchInfo.health_status" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="booker_name">
         <el-input v-model="searchInfo.booker_name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">Search</el-button>
          <el-button icon="refresh" @click="onReset">Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">New</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
              <p style="text-align: center;">Double Check</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">Cancel</el-button>
                <el-button type="primary" @click="onDelete">Confirm</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">Delete</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="Date" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="cluster_name" prop="cluster_name" width="120" />
        <el-table-column align="left" label="pmu_name" prop="pmu_name" width="120" />
        <el-table-column align="left" label="working_status" prop="working_status" width="120" >
          <template #default="scope">
            <el-tag v-if="scope.row.working_status == 'booked'" type="warning" effect="dark">
              booked
            </el-tag>
            <el-tag v-if="scope.row.working_status == 'available'" type="success" effect="dark">
              available
            </el-tag>
            <el-tag v-if="scope.row.working_status == 'working'" type="danger" effect="dark">
              working
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="health_status" prop="health_status" width="120" />
        <el-table-column align="left" label="booker_name" prop="booker_name" width="120" />
        <el-table-column align="left" label="description" prop="description" width="120" />
        <el-table-column align="left" label="Settings">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateClusterFunc(scope.row)">Edit</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">Delete</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="New Cluster">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="cluster_name:"  prop="cluster_name" >
          <el-input v-model="formData.cluster_name" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="pmu_name:"  prop="pmu_name" >
          <el-input v-model="formData.pmu_name" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="working_status:"  prop="working_status" >
            <el-select v-model="formData.working_status" placeholder="请选择" style="width:100%" :clearable="true" >
               <el-option v-for="item in ['available', 'booked', 'working']" :key="item" :label="item" :value="item" />
            </el-select>
        </el-form-item>
        <el-form-item label="health_status:"  prop="health_status" >
          <el-input v-model="formData.health_status" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="description:"  prop="description" >
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">Cancel</el-button>
          <el-button type="primary" @click="enterDialog">Confirm</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Cluster'
}
</script>

<script setup>
import {
  createCluster,
  deleteCluster,
  deleteClusterByIds,
  updateCluster,
  findCluster,
  getClusterList
} from '@/api/clusterStatus'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  cluster_name: '',
  pmu_name: '',
  health_status: '',
  booker_name: '',
  description: '',
  working_status: 'available',
})

// 验证规则
const rule = reactive({
  cluster_name : [{
    required: true,
    message: '',
    trigger: ['input','blur'],
  }],
})

const elFormRef = ref()


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getClusterList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: 'Confirm',
    cancelButtonText: 'Cancel',
    type: 'warning'
  }).then(() => {
    deleteClusterFunc(row)
  })
}


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: 'Please select the data to delete'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.ID)
    })
  const res = await deleteClusterByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Delete Successfully'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateClusterFunc = async(row) => {
  const res = await findCluster({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.recs
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteClusterFunc = async (row) => {
  const res = await deleteCluster({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Delete Successfully'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    cluster_name: '',
    pmu_name: '',
    health_status: '',
    booker_name: '',
    description: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate( async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createCluster(formData.value)
        break
      case 'update':
        res = await updateCluster(formData.value)
        break
      default:
        res = await createCluster(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: 'Create/Edit Successfully'
      })
      closeDialog()
      getTableData()
    }
  })
}
</script>

<style>
</style>
