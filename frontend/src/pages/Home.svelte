<script lang="ts">
    import IconStar from "~icons/tabler/star-filled";
    import QuestionCard from "../lib/components/QuestionCard.svelte";

    import { fetchAllQuestions } from "../lib/utils/fetch.js";
</script>

<div class="">
    <h1 class="text-7xl font-bold">Stuck? Get Unstuck</h1>
    <p class="text-lg">
        Ask product questions, get community answers and move forward
        fast!!!
    </p>
</div>

<hr class="" />

<section class="space-y-16">
    <p class="flex items-center gap-2">
        <IconStar class="text-yellow-500" />Some questions we have so far
    </p>

    {#await fetchAllQuestions()}
        Loading Questions
    {:then questions}
        {#each questions as question, index}
            <QuestionCard {question} />
            {#if index + 1 != questions.length}
                <hr class="text-gray-300">
            {/if}
        {/each}
    {:catch err}
        <p>An error occured: {err}</p>
    {/await}
</section>
