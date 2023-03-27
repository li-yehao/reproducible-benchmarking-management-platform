<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="cluster_name:" prop="cluster_name">
          <el-input v-model="formData.cluster_name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="executor_name:" prop="executor_name">
          <el-input v-model="formData.executor_name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="start_time:" prop="start_time">
          <el-date-picker v-model="formData.start_time" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="end_time:" prop="end_time">
          <el-date-picker v-model="formData.end_time" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="job_status:" prop="job_status">
        <el-select v-model="formData.job_status" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['init', 'working', 'canceled', 'processing', 'process_failed', 'failed', 'done']" :key="item" :label="item" :value="item" />
        </el-select>
        </el-form-item>
        <el-form-item label="cmd_line:" prop="cmd_line">
          <el-input v-model="formData.cmd_line" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="results:" prop="results">
          <el-input v-model="formData.results" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
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
  updateJob_list,
  findJob_list
} from '@/api/jobList'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  cluster_name: '',
  reason: '',
  config: '',
  job_status: '',
  executor_name: '',
  cmd_line: '',
  results: '',
  start_time: new Date(),
  end_time: new Date(),
})
// 验证规则
const rule = reactive({
  cluster_name : [{
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findJob_list({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.rejl
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async() => {
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
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
