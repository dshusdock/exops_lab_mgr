export default () => ({ 
    active: document.getElementsByClassName("page__general")[0],
    parm2: "",
    onMenuClick(event) {
      let text = event.target.innerText;
      console.log("menu clicked: ", text);

      if (text === "General") {
        console.log("General clicked");
        document.getElementsByClassName("page__general")[0].style.display = "flex";
        document.getElementsByClassName("page__dev")[0].style.display = "none";
        document.getElementsByClassName("page__unigy-data-synch")[0].style.display = "none";
      }

      if (text === "Dev") {
        console.log("Dev clicked");      
        document.getElementsByClassName("page__general")[0].style.display = "none";
        document.getElementsByClassName("page__dev")[0].style.display = "flex";
        document.getElementsByClassName("page__unigy-data-synch")[0].style.display = "none";
        
      }

      if (text === "Unigy Data Synch") {
        console.log("Unigy Data Synch clicked");
        document.getElementsByClassName("page__unigy-data-synch")[0].style.display = "flex";
        document.getElementsByClassName("page__general")[0].style.display = "none";
        document.getElementsByClassName("page__dev")[0].style.display = "none";
      }
    },
    onCloseClick(event) {
      console.log("close clicked");
      let modal = document.getElementById("myModal");
      modal.style.display = "none";
    }
})