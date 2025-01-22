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

export function apiGetAlumni(params) {
    return httpInstance({
        url: '/alumni/',
        method: 'get',
        params: params,
    })
}

export function apiCreateAlumni(data) {
    return httpInstance({
        url: '/alumni/',
        method: 'post',
        data: data,
        headers: {
            'Content-Type': 'application/json'
        },
    })
}

export function apiDeleteAlumni(alumni_id) {
    return httpInstance({
        url: `/alumni/${alumni_id}`,
        method: 'delete',
    })
}

export function apiImportAlumni(data) {
    return httpInstance({
        url: '/import/',
        method: 'post',
        data: data,
        headers: {
            'Content-Type': 'multipart/form-data',
        },
    })
}

export function apiGetAlumnusParticipation(alumni_id) {
    return httpInstance({
        url: '/contents/participation?alumni_id=' + alumni_id,
        method: 'get',

    })
}
