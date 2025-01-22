import httpInstance from "@/utils/http.js";

export function apiGetStaffs(search) {
    return httpInstance({
        url: '/staff',
        method: 'get',
        params: {search},
    })
}

export function apiEditStaff(data, staff_id) {
    return httpInstance({
        url: `/staff/${staff_id}`,
        method: 'patch',
        data: data,
        headers: {
            'Content-Type': 'application/json'
        },
    })
}

export function apiCreateStaff(data) {
    return httpInstance({
        url: '/staff',
        method: 'post',
        data: data,
        headers: {
            'Content-Type': 'application/json'
        },
    });
}

export function apiDeleteStaff(staff_id) {
    return httpInstance({
        url: `/staff/${staff_id}`,
        method: 'delete',
    })
}

export function apiResetStaffPassword(staff_id) {
    return httpInstance({
        url: '/staff/reset_password/' + staff_id,
        method: 'post',
    })
}

export function apiChangeStaffPassword(staff_id, data) {
    return httpInstance({
        url: '/staff/change_password/' + staff_id,
        method: 'post',
        data: data,
    })
}
