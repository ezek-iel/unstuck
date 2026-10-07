<script lang="ts">
    import IconChevronUp from "~icons/tabler/chevron-up";
    import type { Question } from "../utils/fetch";

    let { question }: { question: Question } = $props();
</script>

<div class="flex px-8 py-4 items-start gap-8 relative hover:bg-surface-container-high rounded-2xl">
    <div class="space-y-2 border-gray-400 border p-4 rounded-xl">
        <IconChevronUp />
        <div>{question.upvotes}</div>
    </div>
    <div class="space-y-4">
        <p>
            <span class="text-ink-muted px-2 py-1 bg-gray-100"
                >Next.js ecosystem</span
            >
        </p>
        <p class="text-2xl font-bold">
            {question.title}
        </p>

        <p class="text-ink-muted text-lg">
            {question.details}
        </p>
        <ul class="flex flex-wrap gap-2 text-ink-muted capitalize">
            {#if question.tags.length > 0}
                {#each JSON.parse(question.tags) as s}
                    <li class="bg-surface-variant px-3 py-1 rounded">{s}</li>
                {/each}
            {/if}
        </ul>
    </div>

    <a href="/questions/{question.id}" aria-label="Link to {question.title}" class="absolute inset-0 w-full"></a>
</div>
