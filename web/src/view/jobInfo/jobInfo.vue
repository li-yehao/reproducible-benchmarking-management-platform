<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInitInfo" class="demo-form-inline" @keyup.enter="onInitSubmit">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInitInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInitInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="cluster_name">
         <el-input v-model="searchInitInfo.cluster_name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="workload_name">
         <el-input v-model="searchInitInfo.workload_name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="job_status">
         <el-input v-model="searchInitInfo.job_status" placeholder="init" disabled />
        </el-form-item>
            <el-form-item label="svrinfo" prop="svrinfo">
            <el-select v-model="searchInitInfo.svrinfo" clearable placeholder="请选择">
                <el-option
                    key="true"
                    label="true"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="false"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>
            <el-form-item label="emon" prop="emon">
            <el-select v-model="searchInitInfo.emon" clearable placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
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
        <el-form-item>
          <el-button type="primary" icon="search" @click="onInitSubmit">查询</el-button>
          <el-button icon="refresh" @click="onInitReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableInitData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="cluster_name" prop="cluster_name" width="120" />
        <el-table-column align="left" label="workload_name" prop="workload_name" width="130" />
        <el-table-column align="left" label="job_status" prop="job_status" width="120" />
        <el-table-column align="left" label="svrinfo" prop="svrinfo" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.svrinfo) }}</template>
        </el-table-column>
        <el-table-column align="left" label="emon" prop="emon" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.emon) }}</template>
        </el-table-column>
        <el-table-column align="left" label="executor_name" prop="executor_name" width="130" />
         <el-table-column align="left" label="start_time" width="180">
            <template #default="scope">{{ formatDate(scope.row.start_time) }}</template>
         </el-table-column>
         <el-table-column align="left" label="end_time" width="180">
            <template #default="scope">{{ formatDate(scope.row.end_time) }}</template>
         </el-table-column>
        <el-table-column align="left" label="results" prop="results" width="120" />
        <el-table-column align="left" label="log_path" prop="log_path" width="120" />
        <el-table-column align="left" label="svrinfo_path" prop="svrinfo_path" width="120" />
        <el-table-column align="left" label="emon_path" prop="emon_path" width="120" />
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
            :current-page="pageInit"
            :page-size="pageInitSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleInitCurrentChange"
            @size-change="handleInitSizeChange"
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
        <el-form-item label="workload_name">
         <el-input v-model="searchInfo.workload_name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="job_status">
         <el-input v-model="searchInfo.job_status" placeholder="搜索条件" />

        </el-form-item>
            <el-form-item label="svrinfo" prop="svrinfo">
            <el-select v-model="searchInfo.svrinfo" clearable placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>
            <el-form-item label="emon" prop="emon">
            <el-select v-model="searchInfo.emon" clearable placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
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
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
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
        <el-table-column align="left" label="workload_name" prop="workload_name" width="130" />
        <el-table-column align="left" label="job_status" prop="job_status" width="120" />
        <el-table-column align="left" label="svrinfo" prop="svrinfo" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.svrinfo) }}</template>
        </el-table-column>
        <el-table-column align="left" label="emon" prop="emon" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.emon) }}</template>
        </el-table-column>
        <el-table-column align="left" label="executor_name" prop="executor_name" width="130" />
         <el-table-column align="left" label="start_time" width="180">
            <template #default="scope">{{ formatDate(scope.row.start_time) }}</template>
         </el-table-column>
         <el-table-column align="left" label="end_time" width="180">
            <template #default="scope">{{ formatDate(scope.row.end_time) }}</template>
         </el-table-column>
        <el-table-column align="left" label="results" prop="results" width="120" />
        <el-table-column align="left" label="log_path" prop="log_path" width="120" />
        <el-table-column align="left" label="svrinfo_path" prop="svrinfo_path" width="120" />
        <el-table-column align="left" label="emon_path" prop="emon_path" width="120" />
        <el-table-column align="left" label="Settings" width="250">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="cancelJobFunc(scope.row)" v-if="scope.row.job_status === 'working'" :disabled="false">Cancel</el-button>
            <el-button type="primary" link icon="edit" class="table-button" v-else :disabled="true">Cancel</el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateJob_infoFunc(scope.row)">变更</el-button>
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
        <el-form-item label="workload_name:"  prop="workload_name" >
          <el-input v-model="formData.workload_name" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="job_status:"  prop="job_status" >
            <el-select v-model="formData.job_status" placeholder="请选择" style="width:100%" :clearable="true" >
               <el-option v-for="item in ['init', 'canceled', 'working', 'failure', 'done', 'processing', 'process_failure']" :key="item" :label="item" :value="item" />
            </el-select>
        </el-form-item>
        <el-form-item label="svrinfo:"  prop="svrinfo" >
          <el-switch v-model="formData.svrinfo" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="emon:"  prop="emon" >
          <el-switch v-model="formData.emon" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
        <el-form-item label="results:"  prop="results" >
          <el-input v-model="formData.results" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="log_path:"  prop="log_path" >
          <el-input v-model="formData.log_path" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="svrinfo_path:"  prop="svrinfo_path" >
          <el-input v-model="formData.svrinfo_path" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="emon_path:"  prop="emon_path" >
          <el-input v-model="formData.emon_path" :clearable="true"  placeholder="请输入" />
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
  name: 'Job_info'
}
</script>

<script setup>
import {
  createJob_info,
  deleteJob_info,
  deleteJob_infoByIds,
  updateJob_info,
  findJob_info,
  getJob_infoList
} from '@/api/jobInfo'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { findCluster, updateCluster } from '@/api/clusterStatus'
import { findCluster_booking } from '@/api/clusterBooking'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        cluster_name: '',
        workload_name: '',
        svrinfo: false,
        emon: false,
        executor_name: '',
        start_time: new Date(),
        end_time: new Date(),
        results: '',
        log_path: '',
        svrinfo_path: '',
        emon_path: '',
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
  if (searchInfo.value.svrinfo === ""){
      searchInfo.value.svrinfo=null
  }
  if (searchInfo.value.emon === ""){
      searchInfo.value.emon=null
  }
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
  const table = await getJob_infoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  const table = await getJob_infoList({ page: page.value, pageSize: pageSize.value, ...searchInitInfo.value })
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

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteJob_infoFunc(row)
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
  const res = await deleteJob_infoByIds({ ids })
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
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateJob_infoFunc = async(row) => {
  const res = await findJob_info({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reji
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteJob_infoFunc = async (row) => {
  const res = await deleteJob_info({ ID: row.ID })
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
    workload_name: '',
    svrinfo: false,
    emon: false,
    executor_name: '',
    start_time: new Date(),
    end_time: new Date(),
    results: '',
    log_path: '',
    svrinfo_path: '',
    emon_path: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate( async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createJob_info(formData.value)
        break
      case 'update':
        res = await updateJob_info(formData.value)
        break
      default:
        res = await createJob_info(formData.value)
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
    const res0 = await findJob_info({ ID: row.ID })
    // todo: execute ticket really
    const executed = true
    if ( res0.code === 0 && executed ) {
      res0.data.reti.ticket_status = "working"
      const resti = await updateJob_info(res0.data.reti)
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
  
  const res0 = await findJob_info({ ID: row.ID })
  // only the user who create the ticket can cancel a working ticket
  if ( res0.code === 0 && res0.data.reti.user_name === user ) {
    const res1 = await findCluster({ cluster_code: row.cluster_code })
    // todo: cancel ticket really
    const canceled = true
    if ( res1.code === 0 && canceled) {
      res0.data.reti.ticket_status = "canceled"
      const resti = await updateJob_info(res0.data.reti)
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
.el-input {
  width: 100px;
}
</style>
