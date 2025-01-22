import httpInstance from "@/utils/http.js";

export function getStudyHistory() {
    return httpInstance({
        url: '/alumni/studies/get',
        method: 'get',
    })
}

export function editStudy(data, study_id) {
    return httpInstance({
        url: `/alumni/studies/${study_id}`,
        method: 'patch',
        data: data,
        headers: {
            'Content-Type': 'application/json'
        },
    })
}

export function createStudy(data, alumni_id) {
    return httpInstance({
        url: `/alumni/studies/?alumni_id=${alumni_id}`,
        method: 'post',
        data: data,
        headers: {
            'Content-Type': 'application/json'
        },
    });
}

export function apiDeleteStudy(study_id) {
    return httpInstance({
        url: `/alumni/studies/${study_id}`,
        method: 'delete',
    })
}

export function apiGetStudies(alumni_id) {
    return httpInstance({
        url: '/alumni/studies/?alumni_id=' + alumni_id,
        method: 'get',
    })
}
