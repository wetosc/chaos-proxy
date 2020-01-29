<style>
  .menu-column {
    position: relative;
    width: 6%;
    max-width: 50px;
    background: linear-gradient(#2e96ea, #30c8ca);
    min-width: 70px;
  }

  .menu-column .nav {
    list-style-type: none;
    margin: 0 0 100px;
    padding: 0;
  }
  .menu-column .nav li {
    width: 100%;
    height: 3.5em;
    margin: 0;
  }
  .menu-column .nav li:hover {
    background-color: #57a8ed;
    cursor: pointer;
  }
  .menu-column .pulse,
  .menu-column .chat,
  .menu-column .calendar,
  .menu-column .clipboard,
  .menu-column .settings {
    background-image: url("https://svgshare.com/i/4w9.svg");
    background-position: center;
    background-size: 25px;
    background-repeat: no-repeat;
  }
  .menu-column .chat {
    background-image: url("https://svgshare.com/i/4uF.svg");
  }
  .menu-column .calendar {
    background-image: url("https://svgshare.com/i/4vc.svg");
  }
  .menu-column .clipboard {
    background-image: url("https://svgshare.com/i/4wJ.svg");
  }
  .menu-column .settings {
    background-image: url("https://svgshare.com/i/4vd.svg");
  }
</style>


<div class="menu-column">
  <ul class="nav">
    <li class="pulse" />
    <li class="chat" />
    <li class="calendar" />
    <li class="clipboard" />
    <li class="settings" />
  </ul>
</div>
