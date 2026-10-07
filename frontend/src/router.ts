import { createRouter } from "sv-router";

import Home from "./pages/Home.svelte";
import Question from "./pages/Question.svelte";

export const {route} = createRouter({
    "/": Home,
    "/questions/:id": Question
})