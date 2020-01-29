<script>
  // export let name;
  import { token } from "./stores.js";
  import Login from "./Auth/Login.svelte";
  import Main from "./Main/Main.svelte"
</script>

<Main />
<!-- {#if $token}
  <div />
{:else}
  <Login {token} />
{/if} -->
