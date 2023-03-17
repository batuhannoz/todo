import { Pact, Matchers, Verifier } from '@pact-foundation/pact';
import {addTodo} from "./api";
import {describe, beforeAll, afterAll, test, expect} from "vitest";


describe('API Pact test', () => {
    let provider
    beforeAll(async () => {
        provider = new Pact({
            consumer: "consumer-todo",
            provider: "provider-todo",
            port: 3000,
            dir: process.cwd() + '/pacts',
            logLevel: 'debug',
            spec: 2
        })
        await provider.setup()
    })
    afterAll(async () => {
        await provider.finalize()
    })
    describe('sending todo', () => {
        test('send todo', async () => {
            const todo = {
                task: "pliz work this time"
            };
            await provider.addInteraction({
                state: 'sends a todo and expects 200 in return',
                uponReceiving: 'todo takes and saves to database returns 200',
                withRequest: {
                    method: 'POST',
                    path: `/todo/add`
                },
                willRespondWith: {
                    status: 200,
                }
            })

            const result = await addTodo("http://localhost:3000")
            await expect(provider.verify())
        })
    });
});