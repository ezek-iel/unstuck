<script lang="ts">
    import IconChevronUp from "~icons/tabler/chevron-up";
    import IconMessage2Heart from "~icons/tabler/message-2-heart";
    import CommentCard from "../lib/components/CommentCard.svelte";
    import { route } from "../router";
    import { fetchQuestionDetails, getQuestionComments } from "../lib/utils/fetch";

    let currentId = $state(parseInt(route.params.id!) ?? -1);
    let questionDetails = $derived(fetchQuestionDetails(currentId));
    let comments = $derived(getQuestionComments(currentId))
</script>

{#await questionDetails}
    <h1>Loading</h1>
{:then question}
    <div class="space-y-4">
        <p class="text-3xl font-bold text-ink">{question.title}</p>
        <p class="text-lg leading-9">{question.details}</p>

        <div class="flex w-fit items-center space-y-2 border-gray-400 border py-2 px-4 rounded-xl">
            <IconChevronUp />
            <div>{question.upvotes}</div>
        </div>

        <hr class='text-gray-400'>
    </div>

{:catch error}
    <p>An error occured: {error}</p>
{/await}

<div class="space-y-12">
    <h2 class="text-2xl font-bold flex items-center gap-1"><IconMessage2Heart/>Comments</h2>
    {#await comments}
    <h1>Loading Comments</h1>
    {:then data}
        <ul class="space-y-8">
            {#each data as comment}
                <CommentCard {comment}/>
            {/each}
        </ul>
    {:catch err}
        <p>An error occured: {err}</p>
    {/await}
</div>