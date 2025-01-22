import httpInstance from "@/utils/http.js";

export function apiGetContents(search) {
    return httpInstance({
        url: '/contents/default',
        method: 'get',
        params: search
    })
}

export function apiGetContentDetails(content_id) {
    return httpInstance({
        url: '/contents/' + content_id,
        method: 'get',
    })
}

export function apiEditContent(data, content_id) {
    return httpInstance({
        url: `/contents/${content_id}`,
        method: 'patch',
        data: data,
        headers: {
            'Content-Type': 'multipart/form-data'
        },
    })
}

export function apiCreateContent(data) {
    console.log(JSON.stringify(data));
    return httpInstance({
        url: '/contents/',
        method: 'post',
        data: data,
        headers: {
            'Content-Type': 'multipart/form-data'
        },
    });
}

export function apiDeleteContent(content_id) {
    return httpInstance({
        url: `/contents/${content_id}`,
        method: 'delete',
    })
}

export function apiGetEventParticipation({content_id}) {
    return httpInstance({
        url: "/contents/participants",
        method: 'get',
        params: {content_id}
    })
}

export function apiChangeParticipantStatus(data) {
    return httpInstance({
        url: "/contents/change",
        method: 'post',
        data: data,
        headers: {
            'Content-Type': 'application/json',
        },
    })
}

export function apiSendNewsletter(data) {
    return httpInstance({
        url: "/contents/newsletter",
        method: 'post',
        data: data,
        headers: {
            'Content-Type': 'application/json',
        },
    })
}
