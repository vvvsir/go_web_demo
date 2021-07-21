import Layout from '@/layout'

const adminRouter = {
    path: '/admin',
    component: Layout,
    name: 'admin',
    redirect: '/admin/user',
    meta: { title: '后台管理', icon: 'example' },
    alwaysShow: true,
    children: [
        {
            path: 'user',
            name: 'admin-user',
            component: () => import('@/views/admin/user/index'),
            meta: { title: '用户管理', icon: 'user' }
        },
        {
            path: 'role',
            name: 'admin-role',
            component: () => import('@/views/admin/role/index'),
            meta: { title: '角色管理', icon: 'people' }
        },
        {
            path: 'auth',
            name: 'admin-auth',
            component: () => import('@/views/admin/auth/index'),
            meta: { title: '权限管理', icon: 'lock' }
        },
        {
            path: 'user/add',
            name: 'admin-user-add',
            hidden: true,
            component: () => import('@/views/admin/user/add'),
            meta: { title: '添加用户', icon: 'table' }
        },
        {
            path: 'user/edit',
            name: 'admin-user-edit',
            hidden: true,
            component: () => import('@/views/admin/user/edit'),
            meta: { title: '编辑用户', icon: 'table' }
        }
    ]
}

export default adminRouter
