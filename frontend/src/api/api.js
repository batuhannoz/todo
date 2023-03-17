const axios = require('axios');

export const addTodo = (url, todo) => {
    return axios
        .post(url + "/todo/add", todo)
}
