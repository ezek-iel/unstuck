<script lang="ts">
    import type { Comment } from "../utils/fetch";
    let { comment }: { comment: Comment } = $props();
</script>

<li class="text-lg">{comment.details}</li>
