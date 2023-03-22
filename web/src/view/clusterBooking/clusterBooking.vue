<template>
  <div>
    <div class="gva-table-box">
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableDataStatus"
        row-key="ID"
        >
        <el-table-column width="55" />
        <el-table-column align="left" label="cluster_name" prop="cluster_name" width="120" />
        <el-table-column align="left" label="pmu_name" prop="pmu_name" width="120" />
        <el-table-column align="left" label="working_status" prop="working_status" width="120" />
        <el-table-column align="left" label="health_status" prop="health_status" width="120" />
        <el-table-column align="left" label="booker_name" prop="booker_name" width="120" />
        <el-table-column align="left" label="description" prop="description" width="120" />
        <el-table-column align="left" label="Settings" width="280">
          <template #default="scope">
          <el-button type="primary" link icon="edit" class="table-button" @click="book(scope.row)" v-if="scope.row.working_status === 'available'" :disabled="false" >Book</el-button>
          <el-button type="primary" link icon="edit" class="table-button" @click="book(scope.row)" v-else :disabled="true" >Book</el-button>
          <el-button type="primary" link icon="refresh" class="table-button" @click="release(scope.row)" v-if="scope.row.working_status === 'booked'" :disabled="false">Release</el-button>
          <el-button type="primary" link icon="refresh" class="table-button" @click="release(scope.row)" v-else :disabled="true">Release</el-button>
          <!-- <el-button type="primary" link icon="plus" @click="toTarget('Init Job')">Init Job</el-button> -->
          </template>
      </el-table-column>
        <el-table-column width="345" />
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="pageStatus"
            :page-size="pageSizeStatus"
            :page-sizes="[10, 30, 50, 100]"
            :total="totalStatus"
            @current-change="handleCurrentChangeStatus"
            @size-change="handleSizeChangeStatus"
            />
        </div>
    </div>
    <p></p>
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
        <el-form-item label="booker_name">
         <el-input v-model="searchInfo.booker_name" placeholder="搜索条件" />

        </el-form-item>
            <el-form-item label="enabled" prop="enabled">
            <el-select v-model="searchInfo.enabled" clearable placeholder="请选择">
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
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">Search</el-button>
          <el-button icon="refresh" @click="onReset">Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog" disabled>New</el-button>
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
        <el-table-column align="left" label="id" width="80">
            <template #default="scope">{{ scope.row.ID }}</template>
        </el-table-column>
        <el-table-column align="left" label="book_at" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="release_at" width="180">
            <template #default="scope">
              <template v-if="!scope.row.enabled">
                {{ formatDate(scope.row.UpdatedAt) }}
              </template>
            </template>
        </el-table-column>
        <el-table-column align="left" label="cluster_name" prop="cluster_name" width="120" />
        <el-table-column align="left" label="booker_name" prop="booker_name" width="120" />
        <el-table-column align="left" label="enabled" prop="enabled" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.enabled) }}</template>
        </el-table-column>
        <el-table-column align="left" label="booking_reason" prop="booking_reason" width="140" />
        <!-- <el-table-column align="left" label="create_by" width="180">
            <template #default="scope">{{ scope.row.CreatedBy }}</template>
        </el-table-column>
        <el-table-column align="left" label="updated_by" width="180">
            <template #default="scope">{{ scope.row.UpdatedBy }}</template>
        </el-table-column> -->
        <el-table-column align="left" label="Settings">
            <template #default="scope">
            <!-- <el-button type="primary" link icon="edit" class="table-button" @click="updateCluster_bookingFunc(scope.row)">变更</el-button> -->
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="New Cluster Booking">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="cluster_name:"  prop="cluster_name" >
          <el-input v-model="formData.cluster_name" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="booker_name:"  prop="booker_name" >
          <el-input v-model="formData.booker_name" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="enabled:"  prop="enabled" >
          <el-switch v-model="formData.enabled" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="booking_reason:"  prop="booking_reason" >
          <el-input v-model="formData.booking_reason" :clearable="true"  placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">Cancel</el-button>
          <el-button type="primary" @click="enterDialog">Confirm</el-button>
        </div>
      </template>
    </el-dialog>
    <el-dialog v-model="clusterDialogFormVisible" :before-close="closeClusterDialog" title="Input Bookinig Reason">
      <el-form :model="clusterBookingData" label-position="right" ref="elFormRef" :rules="clusterBookingRule" label-width="80px">
        <el-form-item label="booking
        reason:"  prop="booking_reason" >
          <el-input type="textarea" v-model="clusterBookingData.booking_reason" :clearable="true" rows="3"/>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeClusterDialog">Cancel</el-button>
          <el-button type="primary" @click="enterClusterDialog">Confirm</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Cluster_booking'
}
</script>

<script setup>
import {
  createCluster_booking,
  deleteCluster_booking,
  deleteCluster_bookingByIds,
  updateCluster_booking,
  findCluster_booking,
  getCluster_bookingList
} from '@/api/clusterBooking'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { getClusterList, findCluster, updateCluster } from '@/api/clusterStatus'
import { useRouter } from 'vue-router'
import { getUserInfo } from '@/api/user'

const router = useRouter()

// 重定向
const toTarget = (target) => {
  router.push({ name: target })
}

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  cluster_name: '',
  booker_name: '',
  enabled: false,
  booking_reason: '',
})

const clusterBookingData = ref({
  cluster_name: '',
  booker_name: '',
  enabled: false,
  booking_reason: '',
})

const clusterStatusData = ref({
  cluster_name: '',
  working_status: '',
})


// 验证规则
const rule = reactive({
  cluster_name : [{
    required: true,
    message: '',
    trigger: ['input','blur'],
  }],
})

const clusterBookingRule = reactive({
  booking_reason : [{
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

const pageStatus = ref(1)
const totalStatus = ref(0)
const pageSizeStatus = ref(10)
const tableDataStatus = ref([])
const searchInfoStatus = ref({})


// 查询
const getTableDataStatus = async() => {
  const tableStatus = await getClusterList({ page: pageStatus.value, pageSize: pageSizeStatus.value, ...searchInfoStatus.value })
  if (tableStatus.code === 0) {
    tableDataStatus.value = tableStatus.data.list
    totalStatus.value = tableStatus.data.total
    pageStatus.value = tableStatus.data.page
    pageSizeStatus.value = tableStatus.data.pageSize
  }
}

getTableDataStatus()

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.enabled === ""){
      searchInfo.value.enabled=null
  }
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

// 分页
const handleSizeChangeStatus = (val) => {
  pageSizeStatus.value = val
  getTableDataStatus()
}

// 修改页面容量
const handleCurrentChangeStatus = (val) => {
  pageStatus.value = val
  getTableDataStatus()
}

// 查询
const getTableData = async() => {
  const table = await getCluster_bookingList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteCluster_bookingFunc(row)
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
  const res = await deleteCluster_bookingByIds({ ids })
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
const updateCluster_bookingFunc = async(row) => {
  const res = await findCluster_booking({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.recb
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteCluster_bookingFunc = async (row) => {
  const res = await deleteCluster_booking({ ID: row.ID })
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
        booker_name: '',
        enabled: false,
        booking_reason: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate( async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createCluster_booking(formData.value)
        break
      case 'update':
        res = await updateCluster_booking(formData.value)
        break
      default:
        res = await createCluster_booking(formData.value)
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

const clusterDialogFormVisible=ref(false)
// 关闭cluster booking弹窗
const closeClusterDialog = () => {
  clusterDialogFormVisible.value = false
  clusterBookingData.value = {
    cluster_code: '',
    booking_reason: '',
  }
}
// 弹窗确定
const enterClusterDialog = async () => {
  elFormRef.value?.validate( async (valid) => {
    if (!valid) return
    clusterBookingData.value.cluster_name = clusterStatusData.value.cluster_name
    clusterBookingData.value.enabled = true
    const user = await getUserInfo()
    clusterBookingData.value.booker_name = user.data.userInfo.userName
    const resBooking = await createCluster_booking(clusterBookingData.value)
    clusterStatusData.value.working_status = 'booked'
    clusterStatusData.value.booker_name = user.data.userInfo.userName
    const resStatus = await updateCluster(clusterStatusData.value)
    if (resBooking.code === 0 && resStatus.code === 0) {
      ElMessage({
        type: 'success',
        message: 'Booked Successfully '
      })
      closeClusterDialog()
      getTableData()
      getTableDataStatus()
    } else {
      ElMessage({
        type: 'warning',
        message: 'failed'
      })
      return
    }
  })
}

const book = async(row) => {
  const res = await findCluster({ ID: row.ID })
  if (res.code === 0) {
    clusterStatusData.value = res.data.recs
    clusterDialogFormVisible.value = true
  } else {
    ElMessage({
      type: 'warning',
      message: 'cannot find'
    })
    return
  }
}

const searchReleaseInfo = ref({})

const release = async(row) => {
  const rescs = await findCluster({ ID: row.ID })
  clusterStatusData.value = rescs.data.recs
  // todo: double check
  searchReleaseInfo.value.Cluster_name = clusterStatusData.value.cluster_name
  searchReleaseInfo.value.enabled = true
  const rescb = await getCluster_bookingList({ page: page.value, pageSize: pageSize.value, ...searchReleaseInfo.value })
  if (rescs.code === 0 && rescb.code === 0) {
    console.log(rescb.data.list[0])
    clusterBookingData.value = rescb.data.list[0]
    clusterBookingData.value.enabled = false
    console.log(clusterBookingData.value)
    const resBooking = await updateCluster_booking(clusterBookingData.value)
    clusterStatusData.value.working_status = 'available'
    const resStatus = await updateCluster(clusterStatusData.value)
    if (resBooking.code === 0 && resStatus.code === 0) {
      ElMessage({
        type: 'success',
        message: 'Release Successfully'
      }) 
      getTableDataStatus()
      getTableData()
    }
  } else {
    ElMessage({
      type: 'success',
      message: rescb.code.toString(),
    })
  }
}
</script>

<style>
p::before {
  content: "";
  display: block;
  margin: 10px 0;
}
</style>
