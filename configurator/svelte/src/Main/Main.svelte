<script>
  import Menu from "./Menu.svelte";
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
  .content-column .nav .shape {
    font-family: "Nunito", sans-serif;
    font-weight: 200;
    font-size: 6em;
    color: #8b9494;
  }
  .content-column .nav .score {
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
    width: 30%;
    min-width: 240px;
    text-align: center;
    font-size: 1.7em;
    font-weight: 300;
    border-right: 1px solid #dddfe1;
    border-bottom: 1px solid #dddfe1;
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
  .content-column .split-container .value {
    display: block;
    color: #58676c;
    font-size: 2em;
    line-height: 2em;
    width: 180px;
    margin: auto;
    position: relative;
    font-weight: 300;
    margin: 10px auto;
  }

</style>

<div class="container">

  <Menu />

  <div class="content-column">
    <div class="header-container" id="headerContainer">
      <div class="nav">
        <div class="content">
          <p>
            hi
            <span class="name">jana</span>
            , it seems you are in
          </p>
          <span class="shape score">good</span>
          <span class="shape">shape</span>
        </div>
      </div>
    </div>

    <div class="split-container">
      <div class="split">
        <h3>Temperature</h3>
        <div class="value">36.7</div>
      </div>
      <div class="split">
        <h3>Calories burned</h3>
        <div class="value">537</div>
      </div>
      <div class="split">
        <h3>Heart rate</h3>
        <div class="value">87</div>
      </div>
    </div>
  </div>
</div>
