<template>
  <input v-model="task"/>
  <button @click="clicked"/>
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