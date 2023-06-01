<script lang="ts">
  import { Modals, closeModal, openModal } from "svelte-modals";
  import Card from "./Card.svelte";
  import Modal from "./Modal.svelte";

  interface Shrt {
    id: number;
    redirect: string;
    shrt: string;
    random: boolean;
    clicked: number;
  }

  export let shrt: Shrt;
  let showCard = true;
  async function update(data: Shrt): Promise<void> {
    const json = {
      redirect: data.redirect,
      shrt: data.shrt,
      random: data.random,
      id: shrt.id,
      clicked: shrt.clicked,
    };

    await fetch("http://127.0.0.1:8080/shrts", {
      method: "Patch",
      headers: { "content-type": "application/json" },
      body: JSON.stringify(json),
    }).then((res) => {
      console.log(res);
    });
  }

  function handleOpen(shrt: Shrt): void {
    console.log(shrt);
    console.log("Open edit modal");
    openModal(Modal, {
      title: "Update Shrt",
      send: update,
      shrt: shrt.shrt,
      redirect: shrt.redirect,
      random: shrt.random,
    });
  }

  async function deleteShrt(): Promise<void> {
    if (
      confirm("Are you sure you want to delete this Shrt (" + shrt.shrt + ")?")
    ) {
      await fetch("http://127.0.0.1:8080/shrts/" + shrt.id, {
        method: "Delete",
      }).then((res) => {
        showCard = false;
        console.log(res);
      });
    }
  }
</script>

{#if showCard}
  <Card>
    <p>Shrt: http://127.0.0.1:8080/r/{shrt.shrt}</p>
    <p>Redirect: {shrt.redirect}</p>
    <p>Clicked: {shrt.clicked}</p>
    <button class="update" on:click={() => handleOpen(shrt)}>Update</button>
    <button class="delete" on:click={deleteShrt}>Delete</button>
  </Card>
{/if}
<Modals>
  <div
    slot="backdrop"
    class="backdrop"
    on:keydown={closeModal}
    on:click={closeModal}
  />
</Modals>

<style>
  button {
    color: white;
    font-weight: bolder;
    border: none;
    padding: 0.75rem;
    border-radius: 4px;
  }
  .update {
    background-color: yellowgreen;
  }
  .delete {
    background-color: red;
  }
  .backdrop {
    position: fixed;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    background: rgba(255, 255, 255, 0.2);
  }
</style>
