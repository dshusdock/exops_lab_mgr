export default () => ({ 
   someVar: "Hello World",
    // Functions
    onClick(event) {
      console.log("We are in the onClick function" + event.target.id);
      
    },
    onChange(event) {
      console.log("on change clicked" + event.target.id);
      let value = event.target.value;
      let fileName = value.split("\\").pop();
      console.log("fileName: ", fileName);

      document.getElementById("fileChoice").innerText = fileName;
    },
})  