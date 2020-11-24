<script>
  import Menu from "./Menu.svelte";
  import Slider from "./Slider.svelte";
</script>

<style>
  @charset "UTF-8";
  .container {
    font-family: "Roboto", sans-serif;
    background: #7e8b91;
    box-sizing: border-box;
    display: flex;
    display: -webkit-flex;
    flex-direction: row;
    -webkit-flex-direction: row;
    width: 100%;
    min-width: 600px;
    margin: auto;
    justify-content: center;
    min-height: 1200px;
  }

  .content-column {
    background: white;
    width: 100%;
    min-width: 92%;
    z-index: 0;
  }

  .content-column .header-container {
    width: 100%;
    height: 331px;
    overflow: hidden;
  }
  .content-column .nav {
    width: 50%;
    float: left;
    font-size: 0.9em;
    font-weight: 500;
    line-height: 1;
  }
  .content-column .nav .content {
    margin: 50px 0 0 50px;
    white-space: nowrap;
  }
  .content-column .nav p {
    text-transform: uppercase;
    color: #899191;
    letter-spacing: 1px;
  }
  .content-column .nav .title {
    font-family: "Nunito", sans-serif;
    font-weight: 200;
    font-size: 6em;
    color: #8b9494;
  }
  .content-column .nav .title-decoration {
    color: #34d0c4;
  }

  .content-column .split-container {
    display: flex;
    display: -webkit-flex;
    flex-direction: row;
    -webkit-flex-direction: row;
    width: 100%;
    margin: auto;
    justify-content: center;
  }
  .content-column .split-container .split {
    position: relative;
    min-width: 240px;
    text-align: center;
    font-size: 1.7em;
    font-weight: 300;
    border-right: 1px solid #dddfe1;
    border-bottom: 1px solid #dddfe1;
  }

  .col-3 {
    width: 30%;
  }
  .col-2 {
    width: 45%;
  }

  .content-column .split-container .split h3 {
    font-size: 1.1em;
    font-weight: 300;
    margin: 50px 0 0 0;
    color: #566161;
  }
  .content-column .split-container .split:last-child {
    border-right: none;
  }
</style>

<div class="container">

  <!-- <Menu /> -->

  <div class="content-column">
    <div class="header-container" id="headerContainer">
      <div class="nav">
        <div class="content">
          <p>configure your</p>
          <span class="title title-decoration">Chaos</span>
          <span class="title">Proxy</span>
        </div>
      </div>
    </div>

    <div class="split-container">
      <div class="split col-2">
        <h3>Remote server</h3>
      </div>
      <div class="split col-2">
        <h3>Proxy address</h3>
      </div>
    </div>

    <div class="split-container">
      <div class="split col-3">
        <h3>Error 404 Ratio</h3>
        <Slider />
      </div>
      <div class="split col-3">
        <h3>Error 500 Ratio</h3>
        <Slider />
      </div>
      <div class="split col-3">
        <h3>Normal ratio</h3>
        <Slider />
      </div>
    </div>
  </div>
</div>
