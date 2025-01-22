import httpInstance from "@/utils/http.js";

export function apiEditAlumniInfo(data, alumni_id) {
    return httpInstance({
        url: `/alumni/${alumni_id}`,
        method: 'patch',
        data: data,
        headers: {
            'Content-Type': 'application/json'
        },
    })
}

export function apiGetAlumni(search) {
    return httpInstance({
        url: '/alumni/',
        method: 'get',
        params: {search},
    })
}

export function apiResetAlumniPassword(alumni_id) {
    return httpInstance({
        url: '/alumni/reset_password/' + alumni_id,
        method: 'post',
    })
}

export function apiChangeAlumniPassword(alumni_id, data) {
    return httpInstance({
        url: '/alumni/change_password/' + alumni_id,
        method: 'post',
        data: data,
    })
}
