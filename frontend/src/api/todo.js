import axios from "axios";

export const addTodo = (url, task) => {
    axios
        .post(url , {
            task: task
        })
        .then(r => r.data)
}