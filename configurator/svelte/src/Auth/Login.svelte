<script>
  import { fade } from "svelte/transition";
  
  export let token;
  let _token = "";

  let visible = true;
  let formSuccess = "";

  function login() {
    if (!_token) {
      return;
    }
    visible = false;
    formSuccess = "form-success";
    $token = _token;
  }
</script>

<style>
  * {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
    font-weight: 300;
  }
  .wrapper {
    color: white;
    font-weight: 300;

    background: #50a3a2;
    background: linear-gradient(to bottom right, #50a3a2 0%, #53e3a6 100%);
    position: absolute;
    top: 50%;
    left: 0;
    width: 100%;
    height: 400px;
    margin-top: -200px;
    overflow: hidden;
  }
  .wrapper.form-success .container h1 {
    -webkit-transform: translateY(85px);
    transform: translateY(85px);
  }
  .container {
    max-width: 600px;
    margin: 0 auto;
    padding: 80px 0;
    height: 400px;
    text-align: center;
  }
  .container h1 {
    font-size: 40px;
    transition-duration: 1s;
    transition-timing-function: ease-in-put;
    font-weight: 200;
  }
  form {
    padding: 20px 0;
    position: relative;
    z-index: 2;
  }
  form input {
    -webkit-appearance: none;
    -moz-appearance: none;
    appearance: none;
    outline: 0;
    border: 1px solid rgba(255, 255, 255, 0.4);
    background-color: rgba(255, 255, 255, 0.2);
    width: 250px;
    border-radius: 3px;
    padding: 10px 15px;
    margin: 0 auto 10px auto;
    display: block;
    text-align: center;
    font-size: 18px;
    color: white;
    transition-duration: 0.25s;
    font-weight: 300;
  }
  form input:hover {
    background-color: rgba(255, 255, 255, 0.4);
  }
  form input:focus {
    background-color: white;
    width: 300px;
    color: #53e3a6;
  }
  form button {
    -webkit-appearance: none;
    -moz-appearance: none;
    appearance: none;
    outline: 0;
    background-color: white;
    border: 0;
    padding: 10px 15px;
    color: #53e3a6;
    border-radius: 3px;
    width: 250px;
    cursor: pointer;
    font-size: 18px;
    transition-duration: 0.25s;
  }
  form button:hover {
    background-color: #f5f7f9;
  }
</style>

<div class="wrapper {formSuccess}">
  <div class="container">
    <h1>Welcome</h1>

    {#if visible}
      <form class="form" out:fade>
        <input type="text" placeholder="Token" bind:value={_token} />
        <button
          type="submit"
          id="login-button"
          on:click|once|preventDefault={login}>
          Login
        </button>
      </form>
    {/if}
  </div>
</div>
