import request from '@/utils/request'

export function authTree() {
    return request({
        url: '/auth/treemenu'
    })
}
