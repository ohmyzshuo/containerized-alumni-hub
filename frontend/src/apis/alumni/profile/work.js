import httpInstance from "@/utils/http.js";

export function apiGetWorkExperience() {
    return httpInstance({
        url: '/alumni/works/get',
        method: 'get',
    })
}

export function apiEditWork(data, work_id) {
    return httpInstance({
        url: `/alumni/works/${work_id}`,
        method: 'patch',
        data: data,
        headers: {
            'Content-Type': 'application/json'
        },
    })
}

export function apiCreateWork(data, alumni_id) {
    return httpInstance({
        url: `/alumni/works/?alumni_id=${alumni_id}`,
        method: 'post',
        data: data,
        headers: {
            'Content-Type': 'application/json'
        },
    });
}

export function apiDeleteWork(work_id) {
    return httpInstance({
        url: `/alumni/works/${work_id}`,
        method: 'delete',
    })
}

export function apiGetWorks(alumni_id) {
    return httpInstance({
        url: '/alumni/works/?alumni_id=' + alumni_id,
        method: 'get',
    })
}
