import httpInstance from "@/utils/http.js";

export function apiGetPublicationsByToken() {
    return httpInstance({
        url: '/alumni/publications/get',
        method: 'get',
    })
}

export function apiEditPublication(data, publication_id) {
    return httpInstance({
        url: `/alumni/publications/${publication_id}`,
        method: 'patch',
        data: data,
        headers: {
            'Content-Type': 'application/json'
        },
    })
}

export function apiCreatePublication(data, alumni_id) {
    return httpInstance({
        url: `/alumni/publications/?alumni_id=${alumni_id}`,
        method: 'post',
        data: data,
        headers: {
            'Content-Type': 'application/json'
        },
    });
}

export function apiDeletePublication(publication_id) {
    return httpInstance({
        url: `/alumni/publications/${publication_id}`,
        method: 'delete',
    })
}

export function apiGetPublications(alumni_id) {
    return httpInstance({
        url: '/alumni/publications/?alumni_id=' + alumni_id,
        method: 'get',
    })
}

export function apiGetStatistics(params) {
    const queryParams = new URLSearchParams()

    Object.entries(params).forEach(([key, value]) => {
        if (Array.isArray(value)) {
            value.forEach(item => queryParams.append(key, item))
        } else {
            queryParams.append(key, value)
        }
    })
    return httpInstance({
        url: '/alumni/publications/statistics',
        method: 'get',
        params: queryParams,
        paramsSerializer: params => params.toString()
    })
}

export function apiImportPublications(data) {
    return httpInstance({
        url: '/pub-import/',
        method: 'post',
        data: data,
        headers: {
            'Content-Type': 'multipart/form-data',
        },
    })
}