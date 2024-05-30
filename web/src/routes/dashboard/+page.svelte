<script lang="ts">
    import {websocket} from "../authWithCode/stores";
    import type {CommonGuild} from "$lib/guild";
    import CommonGuildSelect from "$lib/commomGuildSelect.svelte"

    const formSubmitAction = process.env.NODE_ENV === "production" ? "/api/uploadAudio" : "http://localhost:3000/api/uploadAudio"
    let ws: any | WebSocket = null;
    let commonGuilds: CommonGuild[] = [];
    let selectedGuild: CommonGuild | null = null;
    let shortAuth: string|null = null;
    let modalOpen = false;
    let selectedFiles: FileList | null = null;
    let sounds: string[] = [];

    const closeModal = () => modalOpen = false;
    const openModal = () => {
        modalOpen = true;
        if (ws !== null) {
            ws.send(JSON.stringify({
                action: "GET_SHORT_AUTH"
            }))
        }
    };

    const onUploadSubmit = async (e: any) => {
        e.preventDefault();
        if (selectedFiles !== null && selectedFiles.length > 0) {
            const formData = new FormData();
            formData.append("audiofile", selectedFiles[0], selectedFiles[0].name);
            await fetch(`${formSubmitAction}?authCode=${shortAuth}`, {
                method: 'POST',
                body: formData
            });
            closeModal();
            if (ws !== null) {
                ws.send(JSON.stringify({
                    action: "GET_ALL_SOUNDS"
                }))
            }
        }
    }

    websocket.subscribe((nws) => {
        ws = nws;
        if (ws !== null) {
            ws.onmessage = (msg: MessageEvent) => {
                const json = JSON.parse(msg.data);
                if (json.action === "GET_COMMON_GUILDS") {
                    commonGuilds = json.content.guilds;
                }
                if (json.action === "GET_SHORT_AUTH") {
                    shortAuth = json.content;
                }
                if (json.action === "GET_ALL_SOUNDS") {
                    sounds = json.content.sounds;
                }
            }
            ws.send(JSON.stringify({
                action: "GET_COMMON_GUILDS"
            }));
            ws.send(JSON.stringify({
                action: "GET_ALL_SOUNDS"
            }));
        }
    });


</script>
<div class={modalOpen ? "modal is-active" : "modal"}>
    <div class="modal-background"></div>
    <div class="modal-card">
        <form on:submit={onUploadSubmit} enctype="multipart/form-data">
            <header class="modal-card-head">
                <p class="modal-card-title">Upload audio</p>
                <button class="delete" aria-label="close" on:click={closeModal}></button>
            </header>
            <section class="modal-card-body">
                <div class="file has-name">
                    <label class="file-label">
                        <input class="file-input" type="file" name="audiofile" accept=".mp3" bind:files={selectedFiles} required />
                        <span class="file-cta">
                          <span class="file-icon">
                            <i class="fas fa-upload"></i>
                          </span>
                          <span class="file-label"> Choose a file… </span>
                        </span>
                        <span class="file-name">
                            {#if selectedFiles}
                                {#each Array.from(selectedFiles) as file}
                                    <p>{file.name}</p>
                                {/each}
                            {/if}
                        </span>
                    </label>
                </div>
            </section>
            <footer class="modal-card-foot">
                <div class="buttons">
                    <button class="button is-success" type="submit">Upload</button>
                    <button class="button" on:click={closeModal}>Cancel</button>
                </div>
            </footer>
        </form>
    </div>
    <button class="modal-close is-large" aria-label="close" on:click={closeModal}></button>
</div>
<div class="fixed-grid has-10-cols">
    <div class="grid">
        <div class="cell is-col-span-2">
            <CommonGuildSelect commonGuilds={commonGuilds} bind:selectedGuild={selectedGuild}/>
        </div>
        <div class="cell is-col-span-7">
            <div class="grid has-3-cols">
                {#each sounds as sound}
                   <div class="cell">
                       <div class="card">
                           <div class="card-content">
                               <div class="content">
                                   {sound}
                               </div>
                           </div>
                       </div>
                   </div>
                {/each}
            </div>
        </div>
        <div class="cell">
            <button class="button is-primary" on:click={openModal}>Upload audio</button>
        </div>
    </div>
</div>