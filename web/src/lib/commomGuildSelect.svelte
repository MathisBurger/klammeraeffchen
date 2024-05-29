<script lang="ts">

    import type {CommonGuild} from "$lib/guild";

    export let commonGuilds: CommonGuild[];
    export let selectedGuild: CommonGuild|null;

    const onClick = (guild: CommonGuild) => () => selectedGuild = guild;
    console.log(selectedGuild);
</script>

<aside class="menu">
    <div class="menu-label">Guilds</div>
    <ul class="menu-list">
        {#each commonGuilds as commonGuild}
            <li class={commonGuild.id === selectedGuild?.id ? "is-active" : ""}><button on:click={onClick(commonGuild)}>{commonGuild.name}</button></li>
        {/each}
    </ul>
</aside>