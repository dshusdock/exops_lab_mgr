export default () => ({ 
    active: document.getElementsByClassName("page__general")[0],
    parm2: "",
    onMenuClick(event) {
      let text = event.target.innerText;
      console.log("menu clicked: ", text);

      if (text === "General") {
        console.log("General clicked");
        document.getElementsByClassName("page__general")[0].style.display = "flex";
        document.getElementsByClassName("page__test")[0].style.display = "none";
        document.getElementsByClassName("page__unigydb")[0].style.display = "none";
      }

      if (text === "Test") {
        console.log("Test clicked");      
        document.getElementsByClassName("page__test")[0].style.display = "flex";
        document.getElementsByClassName("page__general")[0].style.display = "none";
        document.getElementsByClassName("page__unigydb")[0].style.display = "none";
        
      }

      if (text === "Unigy Database") {
        console.log("Unigy Database clicked");
        document.getElementsByClassName("page__unigydb")[0].style.display = "flex";
        document.getElementsByClassName("page__general")[0].style.display = "none";
        document.getElementsByClassName("page__test")[0].style.display = "none";
      }
    },
    onCloseClick(event) {
      console.log("close clicked");
      let modal = document.getElementById("myModal");
      modal.style.display = "none";
    }
})