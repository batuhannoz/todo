import {describe, it, expect} from 'vitest'
import {mount}  from "@vue/test-utils";
import AddTodo from "../AddTodo.vue";

describe('Add Todo Component Test', () => {
    it("Should Render", () => {
        const wrapper = mount(AddTodo)
        expect(wrapper.find("input").exists())
        expect(wrapper.find("button").exists())
    })
})