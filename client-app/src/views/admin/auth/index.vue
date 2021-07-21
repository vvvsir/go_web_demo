<template>
    <div class="app-container">
        <p class="page-title">权限列表</p>
        <el-table
            :data="menusTable"
            style="width: 100%;margin-bottom: 20px;"
            row-key="id"
            border
            :default-expand-all="false"
            :tree-props="{children: 'children', hasChildren: 'hasChildren'}">
            <el-table-column
                prop="id"
                label="编号"
                width="180" />
            <el-table-column
                prop="name"
                label="姓名"
                width="180" />
            <el-table-column
                prop="api"
                label="地址" />
            <el-table-column
                prop="action"
                label="请求方式" />
            <el-table-column
                prop="operation"
                label="操作">
                <template slot-scope="scope">
                    <el-button type="text" size="small" @click="tips">增加</el-button>
                    <!-- 判断下面是否有子菜单，有子菜单不能是有删除按钮 -->
                    <el-button v-if="!scope.row.children" type="text" size="small" @click="tips">删除</el-button>
                    <el-button type="text" size="small" @click="tips">编辑</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>
<script>
import { authTree } from './api'
import { pageMixin } from '@/utils/mixin'
import { removeEmptyChildren } from '@/utils/index'
import { Message } from 'element-ui'

export default {
    name: 'Auth',
    mixins: [pageMixin],
    data() {
        return {
            // 菜单表格结构数据
            menusTable: []
        }
    },
    created() {

    },
    methods: {
        async _getData() {
            let { data } = await authTree()
            data = removeEmptyChildren(data)
            this.menusTable = data
        },
        tips() {
            Message({
                message: '暂未开发',
                type: 'warning',
                duration: 5 * 1000
            })
        }
    }
}
</script>
