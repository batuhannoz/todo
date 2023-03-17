import {describe, it, expect} from 'vitest'
import {mount}  from "@vue/test-utils";
import TodoList from "../TodoList.vue";

describe('Add Todo Component Test', () => {
    it("Should Render", () => {
        const wrapper = mount(TodoList)
        expect(wrapper.find("li").exists())
    })
})