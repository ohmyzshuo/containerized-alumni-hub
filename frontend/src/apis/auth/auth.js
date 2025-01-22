import httpInstance from "@/utils/http.js";


export function apiSendOTP(email) {
    return httpInstance({
        url: '/send-otp?email=' + email,
        method: 'post',
    })
}

export function apiVerifyOTP(email, otp) {
    return httpInstance({
        url: '/send-otp?email=' + email + '&otp=' + otp,
        method: 'post',
    })
}

export function apCheckAlumniExistence(matric_no, otp) {
    return httpInstance({
        url: '/alumni/check?matric_no=' + matric_no,
        method: 'get',
    })
}


export function apiGetByMatricNo(data) {
    return httpInstance({
        url: '/alumni/getByMatricNo',
        method: 'post',
        data: data,
    })
}

export function apiChangeAlumniPassword(alumni_id, data) {
    return httpInstance({
        url: '/alumni/change_password/' + alumni_id,
        method: 'post',
        data: data,
    })
}
