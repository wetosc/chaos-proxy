<script>
  export let value = 50;
</script>

<style>
  .container {
    display: inline-block;
    margin: 15px auto;
    text-align: center;
    width: 100%;
  }

  .slider {
    width: 80%;
    margin: auto;
  }

  input[type="range"] {
    -webkit-appearance: none !important;
    width: 100%;
    height: 8px;
    padding: 2px;
    background-color: #a3cd99;
    border: 1px solid #97c68b;
    border-radius: 8px;
    margin: auto;
    transition: all 0.3s ease;
  }
  input[type="range"]:hover {
    background-color: #b2d5aa;
  }

  input[type="range"]::-webkit-slider-thumb {
    -webkit-appearance: none !important;
    width: 20px;
    height: 20px;
    background-color: #579e81;
    border-radius: 30px;
    box-shadow: 0px 0px 3px #3c6d59;
    transition: all 0.5s ease;
  }
  input[type="range"]::-webkit-slider-thumb:hover {
    background-color: #457d66;
  }
  input[type="range"]::-webkit-slider-thumb:active {
    box-shadow: 0px 0px 1px #3c6d59;
  }

  .value {
    display: block;
    color: #58676c;
    font-size: 2em;
    line-height: 1.5em;
    width: 180px;
    margin: auto;
    position: relative;
    font-weight: 300;
  }
</style>

<div class="container">
  <div class="value">{value}</div>
  <div class="slider">
    <input type="range" min="0" max="100" bind:value />
  </div>
</div>
