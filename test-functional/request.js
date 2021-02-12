import {API_URL,} from './settings'
import axios from 'axios'

export function SendRequest(data, apiTestFile, testFile) {
    let headers = {
        'content-type': 'application/json',
    }
    return axios
        .post(
            API_URL, { query: data},
            {
                headers: headers
            }
        )
        .then(function (response) {
            if (response.data.errors && testFile !== 'mockError')
                console.log('TestFile:', testFile, 'apiTestFile:', apiTestFile, data, ' Error:', response.data.errors)
            return response
        })
}
