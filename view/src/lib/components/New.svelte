<script lang="ts">
  import { openModal } from "svelte-modals";
  import Modal from "./Modal.svelte";

  interface Data {
    redirect: string;
    shrt: string;
    random: boolean;
  }

  async function updateShrt(data: Data): Promise<void> {
    const json = {
      redirect: data.redirect,
      shrt: data.shrt,
      random: data.random,
    };
    console.log(json);
    await fetch("http://127.0.0.1:8080/shrts", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(json),
    }).then((res) => {
      console.log(res);
    });
    console.log(json);
  }

  function handleOpen(): void {
    console.log("Open new modal");
    openModal(Modal, {
      title: "Create New Shrt",
      send: updateShrt,
      redirect: "",
      shrt: "",
      random: false,
    });
  }
</script>

<button on:click={handleOpen}>New</button>

<style>
  button {
    background-color: green;
    color: white;
    font-weight: bold;
    border: none;
    padding: 0.75rem;
    border-radius: 4px;
  }
</style>
