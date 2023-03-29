<template>
  <div>
    <div class="gva-search-box">
      <span style="font-weight: bold; font-size: 16px;">Inited Jobs</span>
      <p></p>
      <el-form :inline="true" :model="searchInitInfo" class="demo-form-inline" @keyup.enter="onInitSubmit">
      
        <el-form-item label="created_at">
          <el-date-picker v-model="searchInitInfo.startCreatedAt" type="datetime" placeholder="start"></el-date-picker>
          —
          <el-date-picker v-model="searchInitInfo.endCreatedAt" type="datetime" placeholder="end"></el-date-picker>
        </el-form-item>

        <el-form-item label="cluster_name">
         <el-input v-model="searchInitInfo.cluster_name" placeholder="criteria" />
        </el-form-item>

        <el-form-item label="executor_name">
         <el-input v-model="searchInitInfo.executor_name" placeholder="criteria" />
        </el-form-item>

        <el-form-item label="start_time">
          <el-date-picker v-model="searchInitInfo.startStart_time" type="datetime" placeholder="start"></el-date-picker>
          —
          <el-date-picker v-model="searchInitInfo.endStart_time" type="datetime" placeholder="end"></el-date-picker>
        </el-form-item>
        <el-form-item label="end_time">
          <el-date-picker v-model="searchInitInfo.startEnd_time" type="datetime" placeholder="start"></el-date-picker>
          —
          <el-date-picker v-model="searchInitInfo.endEnd_time" type="datetime" placeholder="end"></el-date-picker>
        </el-form-item>

        <el-form-item label="job_status">
         <el-input v-model="searchInitInfo.job_status" placeholder="criteria" disabled/>
        </el-form-item>

        <el-form-item label="cmd_line">
         <el-input v-model="searchInitInfo.cmd_line" placeholder="criteria" />
        </el-form-item>

        <el-form-item label="results">
         <el-input v-model="searchInitInfo.results" placeholder="criteria" />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onInitSubmit">Search</el-button>
          <el-button icon="refresh" @click="onInitReset">Reset</el-button>
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
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="reason" prop="reason" width="120" />
        <el-table-column align="left" label="cluster_name" prop="cluster_name" width="120" />
        <el-table-column align="left" label="executor_name" prop="executor_name" width="140" />
        <el-table-column align="left" label="config" prop="config" width="80" >
          <template #default="scope">
            <div>
              <el-popover placement="left-start" trigger="click">
                <div class="popover-box">
                  <pre>{{ scope.row.config }}</pre>
                </div>
                <template #reference>
                  <el-icon style="cursor: pointer;"><warning /></el-icon>
                </template>
              </el-popover>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="job_status" prop="job_status" width="120" >
          <template #default="scope">
            <el-tag v-if="scope.row.job_status == 'INIT'" type="info" effect="dark">
              INIT
            </el-tag>
            <el-tag v-if="scope.row.job_status == 'FAILED'" type="danger" effect="dark">
              FAILED
            </el-tag>
            <el-tag v-if="scope.row.job_status == 'SUCCEEDED'" type="success" effect="dark">
              SUCCEEDED
            </el-tag>
            <el-tag v-if="scope.row.job_status == 'WORKING'" type="warning" effect="dark">
              WORKING
            </el-tag>
            <el-tag v-if="scope.row.job_status == 'CANCELED'" type="danger" effect="dark">
              CANCELED
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="cmd_line" prop="cmd_line" width="100" >
          <template #default="scope">
            <div>
              <el-popover placement="left-start" trigger="click">
                <div class="popover-box">
                  <pre>{{ scope.row.cmd_line }}</pre>
                </div>
                <template #reference>
                  <el-icon style="cursor: pointer;"><warning /></el-icon>
                </template>
              </el-popover>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="results" prop="results" width="80" />
        <el-table-column align="left" label="Settings" width="250">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="executeJobFunc(scope.row)">Execute</el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateJob_listFunc(scope.row)">Edit</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">Delete</el-button>
            </template>
        </el-table-column>
        <el-table-column align="left" label="Date" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="start_time" width="180">
          <template #default="scope">{{ formatDate(scope.row.start_time) }}</template>
        </el-table-column>
        <el-table-column align="left" label="end_time" width="180">
          <template #default="scope">{{ formatDate(scope.row.end_time) }}</template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleInitCurrentChange"
            @size-change="handleInitSizeChange"
            />
        </div>
    </div>
    <p></p>
    <div class="gva-search-box">
      <span style="font-weight: bold; font-size: 16px;">Job List</span>
      <p></p>
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="created_at">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="start"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="end"></el-date-picker>
        </el-form-item>

        <el-form-item label="cluster_name">
         <el-input v-model="searchInfo.cluster_name" placeholder="criteria" />
        </el-form-item>

        <el-form-item label="executor_name">
         <el-input v-model="searchInfo.executor_name" placeholder="criteria" />
        </el-form-item>

        <el-form-item label="start_time">            
          <el-date-picker v-model="searchInfo.startStart_time" type="datetime" placeholder="start"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endStart_time" type="datetime" placeholder="end"></el-date-picker>
        </el-form-item>

        <el-form-item label="end_time">            
          <el-date-picker v-model="searchInfo.startEnd_time" type="datetime" placeholder="start"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endEnd_time" type="datetime" placeholder="end"></el-date-picker>
        </el-form-item>

        <el-form-item label="job_status">
         <el-input v-model="searchInfo.job_status" placeholder="criteria" />
        </el-form-item>

        <el-form-item label="cmd_line">
         <el-input v-model="searchInfo.cmd_line" placeholder="criteria" />
        </el-form-item>

        <el-form-item label="results">
         <el-input v-model="searchInfo.results" placeholder="criteria" />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">Search</el-button>
          <el-button icon="refresh" @click="onReset">Reset</el-button>
        </el-form-item>

      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="search" @click="gotofileserver">FileServer</el-button>
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
        <el-table-column align="left" label="ID" prop="ID" width="80" />
        <el-table-column align="left" label="reason" prop="reason" width="120" />
        <el-table-column align="left" label="cluster_name" prop="cluster_name" width="120" />
        <el-table-column align="left" label="executor_name" prop="executor_name" width="140" />
        <el-table-column align="left" label="config" prop="config" width="80" >
          <template #default="scope">
            <div>
              <el-popover placement="left-start" trigger="click">
                <div class="popover-box">
                  <pre>{{ scope.row.config }}</pre>
                </div>
                <template #reference>
                  <el-icon style="cursor: pointer;"><warning /></el-icon>
                </template>
              </el-popover>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="job_status" prop="job_status" width="120" >
          <template #default="scope">
            <el-tag v-if="scope.row.job_status == 'INIT'" type="info" effect="dark">
              INIT
            </el-tag>
            <el-tag v-if="scope.row.job_status == 'FAILED'" type="danger" effect="dark">
              FAILED
            </el-tag>
            <el-tag v-if="scope.row.job_status == 'SUCCEEDED'" type="success" effect="dark">
              SUCCEEDED
            </el-tag>
            <el-tag v-if="scope.row.job_status == 'WORKING'" type="warning" effect="dark">
              WORKING
            </el-tag>
            <el-tag v-if="scope.row.job_status == 'CANCELED'" type="danger" effect="dark">
              CANCELED
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="cmd_line" prop="cmd_line" width="100" >
          <template #default="scope">
            <div>
              <el-popover placement="left-start" trigger="click">
                <div class="popover-box">
                  <pre>{{ scope.row.cmd_line }}</pre>
                </div>
                <template #reference>
                  <el-icon style="cursor: pointer;"><warning /></el-icon>
                </template>
              </el-popover>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="results" prop="results" width="80" >
          <template #default="scope">
            <div>
              <el-popover placement="left-start" trigger="click">
                <div class="popover-box">
                  <pre>{{ scope.row.results }}</pre>
                </div>
                <template #reference>
                  <el-icon style="cursor: pointer;"><warning /></el-icon>
                </template>
              </el-popover>
            </div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="Settings" width="250">
            <template #default="scope">
            <el-button type="primary" link icon="minus" class="table-button" @click="cancelJobFunc(scope.row)" v-if="scope.row.job_status === 'WORKING'" :disabled="false">Cancel</el-button>
            <el-button type="primary" link icon="minus" class="table-button" v-else :disabled="true">Cancel</el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="openCopyDialog(scope.row)">Copy</el-button>
            <el-button type="primary" link icon="search" @click="gotoEnv(scope.row)">Env</el-button>
            </template>
        </el-table-column>
        <el-table-column align="left" label="Date" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="start_time" width="180">
          <template #default="scope">{{ formatDate(scope.row.start_time) }}</template>
        </el-table-column>
        <el-table-column align="left" label="end_time" width="180">
          <template #default="scope">{{ formatDate(scope.row.end_time) }}</template>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="Init Job">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="cluster:" prop="cluster_name" style="width:80%">
          <el-select v-model="formData.cluster_name" placeholder="select a cluster" >
            <el-option v-for="item in bookedCluster" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="feature:" prop="feature" style="width:80%">
          <el-select v-model="formData.feature" multiple collapse-tags collapse-tags-tooltip placeholder="select features" >
            <el-option v-for="item in featureDict" :key="item.label" :label="item.label" :value="item.label" />
          </el-select>
        </el-form-item>
        <el-form-item label="reason:" prop="reason">
          <el-input type="textarea" v-model="formData.reason" :clearable="true"  placeholder="Reason for this benchmarking" rows="3" />
        </el-form-item>
        <el-form-item label="cmd_line:" prop="cmd_line">
          <el-input type="textarea" v-model="formData.cmd_line" :clearable="true"  placeholder="Pls use absolute path, e.g. python /home/ansible/pkb.py" rows="5" />
        </el-form-item>
        <el-form-item>
          <div>
            <el-button @click="getConfig('gms')">GMS</el-button>
            <el-button @click="getConfig('hotel')">Hotel</el-button>
            <el-button @click="getConfig('social')">Social</el-button>
            <el-button @click="getConfig('socialml')">SocialML</el-button>
            <el-button type="primary" @click="getConfig('clear')">Clear</el-button>
          </div>
          <div>
            <p>Note: Config file will be saved as ~/pkbConfig.yaml on the specified remote host</p>
          </div>
        </el-form-item>
        <el-form-item label="config:" prop="config">
          <el-input type="textarea" v-model="formData.config" :clearable="true"  placeholder="" rows="15" />
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
  executeJob_list,
  cancelJob_list,
  getJob_listList
} from '@/api/jobList'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { getClusterList, updateCluster } from '@/api/clusterStatus'
import { getCluster_bookingList } from '@/api/clusterBooking'
import { getUserInfo } from '@/api/user'
import { createJob_env, updateJob_env, getJob_envList } from '@/api/jobEnv'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  cluster_name: '',
  executor_name: '',
  reason: '',
  config: '',
  start_time: null,
  end_time: null,
  cmd_line: '',
  results: '',
  feature: '',
  job_id: 0,
  key_info: '',
  info_value: '',
})

const featureDict = ref([])
const getFeatureDict = async() => {
  featureDict.value = []
  const res = await getDictFunc("feature")
  for (var i=0; i<res.length; i++) {
    featureDict.value.push({value: res[i].value, label: res[i].label})
  }
}

const getConfig = async(workload) => {
  try {
    if ( workload==="clear" ) {
      formData.value.config = ''
    } else {
      const response = await axios.get('/configs/' + workload + '_config_example.yaml')
      formData.value.config = response.data
    }
  } catch (error) {
    console.log(error);
  }
}

// 验证规则
const rule = reactive({
  cluster_name : [{
    required: true,
    message: '',
    trigger: ['input','blur'],
  }],
  cmd_line : [{
    required: true,
    message: '',
    trigger: ['input','blur'],
  }],
  reason : [{
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

const bookedCluster = ref([])

const getbookedCluster = async() => {
  const user = await getUserInfo()
  const res = await getClusterList({ working_status: "booked", booker_name: user.data.userInfo.userName })
  if (res.code === 0) {
    bookedCluster.value = []
    for (var i=0; i<res.data.list.length; i++) {
      bookedCluster.value.push(res.data.list[i].cluster_name)
    }
  }
}

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
    for (var i=0; i<table.data.list.length; i++) {
      // fetch envs from job_env by job_id
      // todo: too much get env requests
      // const envTable = await getJob_envList({ job_id: table.data.list[i].ID })
      // if (envTable.code === 0) {
      //   for (var j=0; j<envTable.data.list.length; j++) {
      //     const key = envTable.data.list[j].key_info
      //     table.data.list[i][key] = envTable.data.list[j].info_value
      //   }
      // }
    }
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// 查询
const getInitTableData = async() => {
  searchInitInfo.value.job_status = "INIT"
  const table = await getJob_listList({ page: page.value, pageSize: pageSize.value, ...searchInitInfo.value })
  if (table.code === 0) {
    for (var i=0; i<table.data.list.length; i++) {
      // fetch envs from job_env by job_id
      // todo: too much get env requests
      // const envTable = await getJob_envList({ job_id: table.data.list[i].ID })
      // if (envTable.code === 0) {
      //   for (var j=0; j<envTable.data.list.length; j++) {
      //     const key = envTable.data.list[j].key_info
      //     table.data.list[i][key] = envTable.data.list[j].info_value
      //   }
      // }
    }
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
    confirmButtonText: 'Confirm',
    cancelButtonText: 'Cancel',
    type: 'warning'
  }).then(() => {
    deleteJob_listFunc(row)
  })
}

const gotofileserver = () => {
  window.location.href = 'http://10.238.153.58/nfsdata/rbmp/'
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
  const res = await deleteJob_listByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Delete Successfully'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
  }
  getTableData()
  getInitTableData()
}

// 批量删除控制标记
const deleteVisibleInit = ref(false)

// 多选删除
const onDeleteInit = async() => {
  const ids = []
  if (multipleSelectionInit.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: 'Please select the data to delete'
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
      message: 'Delete Successfully'
    })
    if (tableInitData.value.length === ids.length && pageInit.value > 1) {
      pageInit.value--
    }
    deleteVisibleInit.value = false
  }
  getInitTableData()
  getTableData()
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
    getbookedCluster()
    getFeatureDict()
  }
}


// 删除行
const deleteJob_listFunc = async (row) => {
  const res = await deleteJob_list({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Delete Successfully'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
    getInitTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
  getbookedCluster()
  getFeatureDict()
}

// 打开copy弹窗
const openCopyDialog = (row) => {
  type.value = 'create'
  dialogFormVisible.value = true
  formData.value.cmd_line = row.cmd_line
  formData.value.config = row.config
  formData.value.reason = row.reason
  getbookedCluster()
  getFeatureDict()
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    cluster_name: '',
    executor_name: '',
    reason: '',
    config: '',
    start_time: null,
    end_time: null,
    cmd_line: '',
    results: '',
    feature: '',
    job_id: 0,
    key_info: '',
    info_value: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate( async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        const user = await getUserInfo()
        formData.value.executor_name = user.data.userInfo.userName
        formData.value.job_status = "INIT"
        res = await createJob_list(formData.value)
        res = await getJob_listList()
        formData.value.job_id = res.data.list[0].ID
        formData.value.key_info = "feature"
        for (var i=0; i<featureDict.value.length; i++) {
          formData.value.info_value = featureDict.value[i].label
          res = await createJob_env(formData.value)
        }
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
        message: 'Create/Edit Successfully'
      })
      closeDialog()
      getTableData()
      getInitTableData()
    }
  })
}

// execute the inited job
const executeJobFunc = async(row) => {
  const permission = await getCluster_bookingList({ cluster_name: row.cluster_name, enabled: true })
  if ( permission.code !== 0 || permission.data.list.length === 0 ) {
    ElMessage({
      type: 'error',
      message: 'get cluster booking list error'
    })
    return
  }
  const res1 = await getClusterList({ cluster_name: row.cluster_name })
  if ( res1.code !== 0 || res1.data.list.length === 0 ) {
    ElMessage({
      type: 'error',
      message: 'get cluster status list error'
    })
    return
  }
  const user = await getUserInfo()
  if ( user.code !== 0 ) {
    ElMessage({
      type: 'error',
      message: 'get user info error'
    })
    return
  }
  // only the user who booked cluster can execute inited job when the cluster's working status is "booked"
  if ( res1.data.list[0].working_status === "booked" && permission.data.list[0].booker_name === user.data.userInfo.userName ) {
    const res0 = await findJob_list({ ID: row.ID })
    if ( res0.code !== 0 ) {
        ElMessage({
        type: 'error',
        message: 'find job list error'
      })
      return
    }
    // execute job 
    executeJob_list({ ID: row.ID })
    const executed = true
    if ( executed ) {
      res1.data.list[0].working_status = "working"
      const rescs = await updateCluster(res1.data.list[0])
      if ( rescs.code === 0) {
        ElMessage({
          type: 'success',
          message: 'Dispatch job Successfully'
        })
      }
    }
  } else {
    ElMessage({
      type: 'error',
      message: 'No permission for this cluster, please book first'
    })
    return
  } 
  // refresh the web
  router.go(0)
}
// cancel the working job  
const cancelJobFunc = async(row) => {
  // todo: double check
  
  const res0 = await findJob_list({ ID: row.ID })
  // only the user who create the job can cancel a working job  && res0.data.reti.booker_name === user
  if ( res0.code === 0 ) {
    const res1 = await getClusterList({ cluster_name: row.cluster_name })
    if ( res1.code !== 0 ) {
      ElMessage({
        type: 'error',
        message: 'Cannot find cluster'
      })
      return
    }
    // cancel job 
    cancelJob_list({ ID: row.ID })
    const canceled = true
    if ( canceled ) {
      res1.data.list[0].working_status = "booked"
      const rescs = await updateCluster(res1.data.list[0])
      if ( rescs.code === 0 ) {
        ElMessage({
          type: 'success',
          message: 'Cancel Successfully'
        })
        getTableData()
      }
    }
  } else {
    ElMessage({
      type: 'error',
      message: 'Cannot find job'
    })
  }
}

const gotoEnv = (row) => {
  if (row) {
    router.push({ name: 'jobEnv', params: {
      job_id: row.ID
    }})
  } else {
    router.push({ name: 'jobEnv' })
  }
}
</script>

<style>
.popover-box {
  background: #112435;
  color: #f08047;
  height: 600px;
  width: 420px;
  overflow: auto;
  overflow-y: auto;
}
.popover-box::-webkit-scrollbar {
  display: none; /* Chrome Safari */
}
.popover-box pre {
  white-space: pre-wrap;
}
p::before {
  content: "";
  display: block;
  margin: 15px 0;
}
</style>