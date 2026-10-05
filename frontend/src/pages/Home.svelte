<script lang="ts">
    import IconChevronUp from "~icons/tabler/chevron-up";
    import IconStar from "~icons/tabler/star-filled";
</script>

<div class="container mx-auto my-20 text-ink max-w-4xl space-y-8">
    <div class="">
        <h1 class="text-7xl font-bold">Stuck? Get Unstuck</h1>
        <p class="text-lg">
            Ask product questions, get community answers and move forward fast!!!
        </p>
    </div>

    <hr class="" />

    <section class="space-y-16">
        <p class="flex items-center gap-2"><IconStar class="text-yellow-500"/>Some questions we have so far</p>
        {#each { length: 9 } as _,i}
            <div class="flex items-start gap-8">
                <div class="space-y-2 border-gray-400 border p-4 rounded-xl">
                    <IconChevronUp />
                    <div>20</div>
                </div>
                <div class="space-y-4">
                    <p><span class="text-ink-muted px-2 py-1 bg-gray-100">Next.js ecosystem</span></p>
                    <p class="text-2xl font-bold">
                        How to prevent Next.js 15 Server Action race condition
                        on optimistic mutations?
                    </p>

                    <p class="text-ink-muted text-lg">
                        When triggering consecutive fast updates using
                        `useOptimistic` alongside `useActionState`, the stale
                        promise response overwrites the most recent optimistic
                        state. Is there an idiomatic abort controller pattern
                        for server actions?
                    </p>
                    <ul class="flex flex-wrap gap-2 text-ink-muted capitalize">
                    {#each ["nextjs", "react"] as s}
                        <li class="bg-surface-variant px-3 py-1 rounded">{s}</li>
                        {/each}
                    </ul>
                    {#if i != 8} 
                        <hr class="text-gray-400">
                    {/if}
                </div>
            </div>

        {/each}
    </section>
</div>
