<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onInitSubmit">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInitInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInitInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="cluster_name">
         <el-input v-model="searchInitInfo.cluster_name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="executor_name">
         <el-input v-model="searchInitInfo.executor_name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="start_time">
            
            <el-date-picker v-model="searchInitInfo.startStart_time" type="datetime" placeholder="搜索条件（起）"></el-date-picker>
            —
            <el-date-picker v-model="searchInitInfo.endStart_time" type="datetime" placeholder="搜索条件（止）"></el-date-picker>

        </el-form-item>
        <el-form-item label="end_time">
            
            <el-date-picker v-model="searchInitInfo.startEnd_time" type="datetime" placeholder="搜索条件（起）"></el-date-picker>
            —
            <el-date-picker v-model="searchInitInfo.endEnd_time" type="datetime" placeholder="搜索条件（止）"></el-date-picker>

        </el-form-item>
        <el-form-item label="job_status">
         <el-input v-model="searchInitInfo.job_status" placeholder="搜索条件" disabled/>

        </el-form-item>
        <el-form-item label="cmd_line">
         <el-input v-model="searchInitInfo.cmd_line" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="results">
         <el-input v-model="searchInitInfo.results" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onInitSubmit">查询</el-button>
          <el-button icon="refresh" @click="onInitReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">New</el-button>
            <el-popover v-model:visible="deleteVisibleInit" placement="top" width="180">
            <p style="text-align: center;">Double Check</p>
            <div style="text-align: center; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisibleInit = false">Cancel</el-button>
                <el-button type="primary" @click="onDeleteInit">Confirm</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelectionInit.length" @click="deleteVisibleInit = true">Delete</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableInitData"
        row-key="ID"
        @selection-change="handleSelectionChangeInit"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="cluster_name" prop="cluster_name" width="120" />
        <el-table-column align="left" label="executor_name" prop="executor_name" width="120" />
         <el-table-column align="left" label="start_time" width="180">
            <template #default="scope">{{ formatDate(scope.row.start_time) }}</template>
         </el-table-column>
         <el-table-column align="left" label="end_time" width="180">
            <template #default="scope">{{ formatDate(scope.row.end_time) }}</template>
         </el-table-column>
        <el-table-column align="left" label="job_status" prop="job_status" width="120" />
        <el-table-column align="left" label="cmd_line" prop="cmd_line" width="120" />
        <el-table-column align="left" label="results" prop="results" width="120" />
        <el-table-column align="left" label="Settings" width="250">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="executeJobFunc(scope.row)">Execute</el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateJob_infoFunc(scope.row)">Edit</el-button>
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
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="cluster_name">
         <el-input v-model="searchInfo.cluster_name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="executor_name">
         <el-input v-model="searchInfo.executor_name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="start_time">
            
            <el-date-picker v-model="searchInfo.startStart_time" type="datetime" placeholder="搜索条件（起）"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endStart_time" type="datetime" placeholder="搜索条件（止）"></el-date-picker>

        </el-form-item>
        <el-form-item label="end_time">
            
            <el-date-picker v-model="searchInfo.startEnd_time" type="datetime" placeholder="搜索条件（起）"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endEnd_time" type="datetime" placeholder="搜索条件（止）"></el-date-picker>

        </el-form-item>
        <el-form-item label="job_status">
         <el-input v-model="searchInfo.job_status" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="cmd_line">
         <el-input v-model="searchInfo.cmd_line" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="results">
         <el-input v-model="searchInfo.results" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">New</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="180">
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
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="cluster_name" prop="cluster_name" width="120" />
        <el-table-column align="left" label="executor_name" prop="executor_name" width="120" />
         <el-table-column align="left" label="start_time" width="180">
            <template #default="scope">{{ formatDate(scope.row.start_time) }}</template>
         </el-table-column>
         <el-table-column align="left" label="end_time" width="180">
            <template #default="scope">{{ formatDate(scope.row.end_time) }}</template>
         </el-table-column>
        <el-table-column align="left" label="job_status" prop="job_status" width="120" />
        <el-table-column align="left" label="cmd_line" prop="cmd_line" width="120" />
        <el-table-column align="left" label="results" prop="results" width="120" />
        <el-table-column align="left" label="Settings" width="250">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="cancelJobFunc(scope.row)" v-if="scope.row.job_status === 'working'" :disabled="false">Cancel</el-button>
            <el-button type="primary" link icon="edit" class="table-button" v-else :disabled="true">Cancel</el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateJob_listFunc(scope.row)">Edit</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="cluster_name:"  prop="cluster_name" >
          <el-input v-model="formData.cluster_name" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="executor_name:"  prop="executor_name" >
          <el-input v-model="formData.executor_name" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="start_time:"  prop="start_time" >
          <el-date-picker v-model="formData.start_time" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="end_time:"  prop="end_time" >
          <el-date-picker v-model="formData.end_time" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="job_status:"  prop="job_status" >
            <el-select v-model="formData.job_status" placeholder="请选择" style="width:100%" :clearable="true" >
               <el-option v-for="item in ['init', 'working', 'canceled', 'processing', 'process_failed', 'failed', 'done']" :key="item" :label="item" :value="item" />
            </el-select>
        </el-form-item>
        <el-form-item label="cmd_line:"  prop="cmd_line" >
          <el-input v-model="formData.cmd_line" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="results:"  prop="results" >
          <el-input v-model="formData.results" :clearable="true"  placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Job_list'
}
</script>

<script setup>
import {
  createJob_list,
  deleteJob_list,
  deleteJob_listByIds,
  updateJob_list,
  findJob_list,
  getJob_listList
} from '@/api/jobList'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { findCluster, updateCluster } from '@/api/clusterStatus'
import { findCluster_booking } from '@/api/clusterBooking'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  cluster_name: '',
  executor_name: '',
  start_time: new Date(),
  end_time: new Date(),
  cmd_line: '',
  results: '',
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

const pageInit = ref(1)
const totalInit = ref(0)
const pageInitSize = ref(10)
const tableInitData = ref([])
const searchInitInfo = ref({})

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

// 重置
const onInitReset = () => {
  searchInitInfo.value = {}
  getInitTableData()
}
// 搜索
const onInitSubmit = () => {
  pageInit.value = 1
  pageInitSize.value = 10
  if (searchInitInfo.value.svrinfo === ""){
    searchInitInfo.value.svrinfo=null
  }
  if (searchInitInfo.value.emon === ""){
    searchInitInfo.value.emon=null
  }
  getInitTableData()
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

// 分页
const handleInitSizeChange = (val) => {
  pageInitSize.value = val
  getInitTableData()
}

// 修改页面容量
const handleInitCurrentChange = (val) => {
  pageInit.value = val
  getInitTableData()
}

// 查询
const getTableData = async() => {
  const table = await getJob_listList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// 查询
const getInitTableData = async() => {
  searchInitInfo.value.job_status = "init"
  const table = await getJob_listList({ page: page.value, pageSize: pageSize.value, ...searchInitInfo.value })
  if (table.code === 0) {
    tableInitData.value = table.data.list
    totalInit.value = table.data.total
    pageInit.value = table.data.page
    pageInitSize.value = table.data.pageSize
  }
}
getInitTableData()

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

// 多选数据
const multipleSelectionInit = ref([])
// 多选
const handleSelectionChangeInit  = (val) => {
    multipleSelectionInit.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteJob_listFunc(row)
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
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.ID)
    })
  const res = await deleteJob_listByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
    getInitTableData()
  }
}

// 批量删除控制标记
const deleteVisibleInit = ref(false)

// 多选删除
const onDeleteInit = async() => {
  const ids = []
  if (multipleSelectionInit.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelectionInit.value &&
    multipleSelectionInit.value.map(item => {
      ids.push(item.ID)
    })
  const res = await deleteJob_listByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableInitData.value.length === ids.length && pageInit.value > 1) {
      pageInit.value--
    }
    deleteVisibleInit.value = false
    getInitTableData()
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateJob_listFunc = async(row) => {
    const res = await findJob_list({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rejl
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteJob_listFunc = async (row) => {
    const res = await deleteJob_list({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
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
        executor_name: '',
        start_time: new Date(),
        end_time: new Date(),
        cmd_line: '',
        results: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createJob_list(formData.value)
                  break
                case 'update':
                  res = await updateJob_list(formData.value)
                  break
                default:
                  res = await createJob_list(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}

// execute the inited ticket
const executeJobFunc = async(rou) => {
  // todo: double check
  const permission = await findCluster_booking({ cluster_code: row.cluster_code, enabled: true })
  const res1 = await findCluster({ cluster_code: row.cluster_code })
  // only the user who booked cluster can execute inited ticket when the cluster's working status is "booked"
  if ( permission.code === 0 && res1.code === 0 && res1.data.recs.working_status === "booked" && permission.data.recb.user_name === user ) {
    const res0 = await findJob_list({ ID: row.ID })
    // todo: execute ticket really
    const executed = true
    if ( res0.code === 0 && executed ) {
      res0.data.reti.ticket_status = "working"
      const resti = await updateJob_list(res0.data.reti)
      res1.data.recs.working_status = "working"
      const rescs = await updateCluster(res1.data.recs)
      if (resti.code === 0 && rescs.code === 0) {
        ElMessage({
          type: 'success',
          message: 'Execute Successfully'
        })
      }
    }
  }
  
}
// cancel the working ticket  
const cancelJobFunc = async(row) => {
  // todo: double check
  
  const res0 = await findJob_list({ ID: row.ID })
  // only the user who create the ticket can cancel a working ticket
  if ( res0.code === 0 && res0.data.reti.user_name === user ) {
    const res1 = await findCluster({ cluster_code: row.cluster_code })
    // todo: cancel ticket really
    const canceled = true
    if ( res1.code === 0 && canceled) {
      res0.data.reti.ticket_status = "canceled"
      const resti = await updateJob_list(res0.data.reti)
      res1.data.reti.working_status = "booked"
      const rescs = await updateCluster(res1.data.reti)
      if (resti.code === 0 && rescs.code === 0) {
        ElMessage({
          type: 'success',
          message: 'Cancel Successfully'
        })
      }
    }
  }
}
</script>

<style>
</style>
