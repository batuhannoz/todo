<template>
  <input v-model="task"/>
  <button @click="clicked" style="width: 100px;margin-left: 20px">Add Task</button>
</template>

<script>
export default {
  emits: ["addTask"],
  data() {
    return {
      task: ""
    }
  },
  methods: {
    clicked() {
      this.$emit('addTask', this.task)
      this.task = ""
    }
  }
}
</script>