export default () => ({ 
    previousTarget: null,
    previousTarget2: null,
    // Functions
    onClick(event) {
      console.log("We are in the onClick function - " + event.target.innerText);
      if (this.previousTarget) {
        this.previousTarget.style.backgroundColor = "#0964b0";
        this.previousTarget2.style.backgroundColor = "#0964b0";
      }
      this.previousTarget = event.target;
      event.target.style.backgroundColor = "#40A9BF";
      
    },
    onClick2(event) {
      console.log("We are in the onClick2 function - " + event.target.innerText);
      if (this.previousTarget2) {
        this.previousTarget2.style.backgroundColor = "#0964b0";
      }
      this.previousTarget2 = event.target;
      event.target.style.backgroundColor = "#40A9BF";
      
    },
    onChange(event) {
      console.log("on change clicked" + event.target.innerText);
      let value = event.target.value;
      let fileName = value.split("\\").pop();
      console.log("fileName: ", fileName);

      document.getElementById("fileChoice").innerText = fileName;
    },
})  