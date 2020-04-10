<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
   
        <el-form-item label="昵称">
          <el-input placeholder="昵称/所在地(模糊查询)" v-model="searchInfo.name"></el-input>
        </el-form-item>
         <el-form-item label="开播状态">
               <el-select placeholder="请选择" clearable v-model="searchInfo.broadcast">
            <el-option
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
              v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="审核状态">
               <el-select placeholder="请选择" clearable v-model="searchInfo.audit">
            <el-option
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
              v-for="item in auditOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
       
      </el-form>
    </div>
    <el-table :data="tableData" border stripe>
      <el-table-column label="id" min-width="60" prop="ID"></el-table-column>
      <!--<el-table-column label="主播头像" min-width="150" prop="logoImg"></el-table-column>-->
      <el-table-column label="主播头像" min-width="50">
        <template slot-scope="scope">
          <div :style="{'textAlign':'center'}">
            <img :src="scope.row.logoImg" height="50" width="50" />
          </div>
        </template>
      </el-table-column>
      
     <!-- <el-table-column label="直播间图" min-width="150" prop="gapImg"></el-table-column>-->
       <el-table-column label="直播间图" min-width="50">
        <template slot-scope="scope">
          <div :style="{'textAlign':'center'}">
            <img :src="scope.row.gapImg" height="100" width="100" />
          </div>
        </template>
      </el-table-column>
      
        <!-- <el-table-column label="个人空间首图" min-width="150" prop="zoneImg"></el-table-column>-->

       <el-table-column label="个人空间首图" min-width="50">
        <template slot-scope="scope">
          <div :style="{'textAlign':'center'}">
            <img :src="scope.row.zoneImg" height="100" width="100" />
          </div>
        </template>
      </el-table-column>
      

      <el-table-column label="关键描述" min-width="150" prop="description"></el-table-column>
      <!-- <el-table-column label="状态" min-width="150" prop="state"></el-table-column>-->


      <el-table-column label="状态" min-width="150" prop="method">
        <template slot-scope="scope">
          <div>
           <!-- {{scope.row.state}}-->
            <el-tag
              :key="scope.row.methodFiletr"
              :type="scope.row.broadcast|tagTypeFiletr"
              effect="dark"
              size="mini"
            >{{scope.row.broadcast|methodFiletr}}</el-tag>
            <!-- {{scope.row.method|methodFiletr}} -->
          </div>
        </template>
      </el-table-column>

    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增Api">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="apiForm">
        <el-form-item label="路径" prop="path">
          <el-input autocomplete="off" v-model="form.path"></el-input>
        </el-form-item>
        <el-form-item label="请求" prop="method">
          <el-select placeholder="请选择" v-model="form.method">
            <el-option
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
              v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="api分组" prop="group">
          <el-input autocomplete="off" v-model="form.group"></el-input>
        </el-form-item>
        <el-form-item label="api简介" prop="description">
          <el-input autocomplete="off" v-model="form.description"></el-input>
        </el-form-item>
      </el-form>
      <div class="warning">新增Api需要在角色管理内配置权限才可使用</div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

import {

  getAnchorList

} from '@/api/anchor'
import infoList from '@/components/mixins/infoList'


const stateOptions = [
  {
    value: '-1',
    label: '不可开播',
    type: 'danger'
  },
  {
    value: '0',
    label: '可开播',
    type: 'warning'
  },
  
  {
    value: '1',
    label: '开播中',
    type: 'success'
  }
]

const methodOptions = [
  {
    value: 'All',
    label: '所有',
    type: ''
  },
  {
    value: 'Disable',
    label: '不可开播',
    type: 'danger'
  },
  {
    value: 'Pending',
    label: '可开播',
    type: 'warning'
  },
  {
    value: 'Online',
    label: '开播中',
    type: 'success'
  }
]

const auditOptions = [
  {
    value: 'All',
    label: '全部',
    type: 'success'
  },
  {
    value: 'Success',
    label: '通过',
    type: ''
  },
  {
    value: 'Failure',
    label: '不通过',
    type: ''
  }
]
export default {
  name: 'Api',
  mixins: [infoList],
  data() {
    return {
      listApi: getAnchorList,
      listKey: 'list',
      dialogFormVisible: false,
      form: {
        path: '',
        group: '',
        method: '',
        description: ''
      },
      methodOptions: methodOptions,
      auditOptions: auditOptions,
      stateOptions: stateOptions,
      type: '',
      rules: {
        path: [{ required: true, message: '请输入api路径', trigger: 'blur' }],
        group: [{ required: true, message: '请输入组名称', trigger: 'blur' }],
        method: [
          { required: true, message: '请选择请求方式', trigger: 'blur' }
        ],
        description: [
          { required: true, message: '请输入api介绍', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    initForm() {
      this.form = {
        path: '',
        group: '',
        description: '',
        method: ''
      }
    },
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    openDialog(type) {
      this.type = type
      this.dialogFormVisible = true
    },
   
    async enterDialog() {
      this.$refs.apiForm.validate(async valid => {
        if (valid) {
          switch (this.type) {
            case 'addApi':
              {
                // const res = await createApi(this.form)
                // if (res.success) {
                  this.$message({
                    type: 'success',
                    message: '添加成功',
                    showClose: true
                  })
                // }
                this.getTableData()
                this.closeDialog()
              }

              break
            case 'edit':
              {
                // const res = await updataApi(this.form)
                // if (res.success) {
                  this.$message({
                    type: 'success',
                    message: '添加成功',
                    showClose: true
                  })
                // }
                this.getTableData()
                this.closeDialog()
              }
              break
            default:
              {
                this.$message({
                  type: 'error',
                  message: '未知操作',
                  showClose: true
                })
              }
              break
          }
        }
      })
    }
  },
  filters: {
    methodFiletr(value) {
      const target = stateOptions.filter(item => item.value === value)[0]
      // return target && `${target.label}(${target.value})`
      return target && `${target.label}`
    },
    tagTypeFiletr(value) {
      const target = stateOptions.filter(item => item.value === value)[0]
      return target && `${target.type}`
    }
  }
}
</script>
<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.el-tag--mini {
  margin-left: 5px;
}
.warning {
  color: #dc143c;
}
</style>