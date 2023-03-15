<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="cluster_name:" prop="cluster_name">
          <el-input v-model="formData.cluster_name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="workload_name:" prop="workload_name">
          <el-input v-model="formData.workload_name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="job_status:" prop="job_status">
        <el-select v-model="formData.job_status" placeholder="请选择" style="width:100%" :clearable="true">
          <el-option v-for="item in ['init', 'canceled', 'working', 'failure', 'done', 'processing', 'process_failure']" :key="item" :label="item" :value="item" />
        </el-select>
        </el-form-item>
        <el-form-item label="svrinfo:" prop="svrinfo">
          <el-switch v-model="formData.svrinfo" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="emon:" prop="emon">
          <el-switch v-model="formData.emon" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
        <el-form-item label="results:" prop="results">
          <el-input v-model="formData.results" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="log_path:" prop="log_path">
          <el-input v-model="formData.log_path" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="svrinfo_path:" prop="svrinfo_path">
          <el-input v-model="formData.svrinfo_path" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="emon_path:" prop="emon_path">
          <el-input v-model="formData.emon_path" :clearable="true" placeholder="请输入" />
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
  name: 'Job_info'
}
</script>

<script setup>
import {
  createJob_info,
  updateJob_info,
  findJob_info
} from '@/api/jobInfo'

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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findJob_info({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reji
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
