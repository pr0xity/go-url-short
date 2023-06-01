<script lang="ts">
  import { closeModal } from "svelte-modals";
  export let title: string;
  export let isOpen: boolean;
  export let redirect: string;
  export let shrt: string;
  export let random: boolean;
  export let send: (data: Data) => void;

  interface Data {
    redirect: string;
    shrt: string;
    random: boolean;
  }

  let data: Data = {
    redirect: redirect,
    shrt: shrt,
    random: random,
  };
</script>

{#if isOpen}
  <div role="dialog" class="modal">
    <div class="contents">
      <h2>{title}</h2>
      <label for="redirect">Redirect to:</label>
      <input
        type="text"
        name="redirect"
        id="redirect"
        bind:value={data.redirect}
      />

      <label for="shrt">Shrt</label>
      <input
        type="text"
        name="shrt"
        id="shrt"
        bind:value={data.shrt}
        disabled={data.random}
        class={data.random ? "disabled" : ""}
      />

      <label for="random">Randomly generate shrt</label>
      <input
        type="checkbox"
        name="random"
        id="random"
        bind:checked={data.random}
      />

      <div class="actions">
        <button on:click={closeModal}>Cancel</button>
        <button on:click={() => send(data)} on:click={closeModal}>Save</button>
      </div>
    </div>
  </div>
{/if}

<style>
  .modal {
    color: black;
    position: fixed;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    display: flex;
    justify-content: center;
    align-items: center;

    /* allow click-through to backdrop */
    pointer-events: none;
  }

  .contents {
    min-width: 500px;
    border-radius: 6px;
    padding: 16px;
    background: white;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    pointer-events: auto;
  }

  h2 {
    text-align: center;
    font-size: 24px;
  }

  .actions {
    margin-top: 32px;
    display: flex;
    justify-content: space-between;
    gap: 8px;
  }

  .disabled {
    background: slategray;
  }
</style>
