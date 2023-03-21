<template>
  <AddTodo @addTask="addTask" />
  <TodoList :todos="todos"/>
</template>

<script>
import AddTodo from "./components/AddTodo.vue";
import TodoList from "./components/TodoList.vue";
import {addTodo} from "./api/api";

export default {
  components: { AddTodo, TodoList },
  data() {
    return {
      todos: []
    };
  },
  methods: {
    addTask(task) {
      addTodo(process.env.API_HOST + ":" + process.env.API_PORT, {task: task})
          .then((r) => {
            if (r.status === 200) {
              this.todos.push({
                "task": task
              })
            }
          })
    }
  }
}
</script>

<style scoped>

</style>
