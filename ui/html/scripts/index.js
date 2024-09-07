import Alpine from 'alpinejs'
import myData from './mydata'
import modalData from './modaldata'
import hdrData from './hdrdata'
import settings from './settings'
import sidenav from './sidenav'
import lstable from './lstable'
import login from './login'

Alpine.data('myData', myData)
Alpine.data('modalData', modalData)
Alpine.data('hdrData', hdrData)
Alpine.data('settings', settings)
Alpine.data('sidenav', sidenav)
Alpine.data('lstable', lstable)
Alpine.data('login', login)

Alpine.start()

console.log("Release the hounds!!......again!");


// document.addEventListener("alpine:init", () => {
//     Alpine.data('myData', () => ({
//     target: "testing...",
//     flag: true,
//     drop: false,
//     win1: true,
//     win2: false,
//     onClick() {
//       console.log("clicked -- flag: ", this.flag);
//       let el = document.getElementById("test_form");
//       if (this.flag) {
//         // el.className += " info-form--on"
//         el.classList.add("info-form--on");
//         this.flag = false;
//       } else {
//         el.classList.remove("info-form--on");
//         el.classList.add("info-form--off");

//         // el.className += " info-form--off"
//         this.flag = true;
//       }
//     },
//     handleTabClick(event) {
//       console.log("clicked tab - " + event.target.innerText);

//       let el = event.target;
//       let clickedText = el.innerText;
//       let parent = event.currentTarget;

//       let nodes = parent.childNodes;
//       console.log("nodes: ", nodes);

//       nodes.forEach((element) => {
//         if (element.tagName === "SPAN") {
//           console.log("element: ", element.innerText);

//           if (clickedText === element.innerText) {
//             console.log("found it: ", element);
//             element.style.backgroundColor = "blue";
//           } else {
//             console.log("found other: ", element.tagName);
//             element.style.backgroundColor = "gray";
//           }

//           if (clickedText === "Test1") {
//             this.win1 = true;
//             this.win2 = false;
//           } else {
//             this.win1 = false;
//             this.win2 = true;
//           }
//         }
//       });
//     },
//     onSettingClick(event) {
//       this.drop = !this.drop;
//     },
//     testThis() {
//       console.log("testThis: ", this);
//       this.drop = false;
//     },
//     testThis2(event) {
//       console.log("Got the focus");
//     },
//     })),
//     Alpine.data('modalData', () => ({
//     btn: document.getElementById("myBtn"),
//     span: document.getElementsByClassName("close")[0],
//     // Functions
//     onFileClick(event) {
//       console.log("on file clicked");
//       document.getElementById("fileInput").click();
//     },
//     onChange(event) {
//       console.log("on change clicked" + event.target.value);
//       let value = event.target.value;
//       let fileName = value.split("\\").pop();
//       console.log("fileName: ", fileName);

//       document.getElementById("fileChoice").innerText = fileName;
//     },
//     onCloseClick(event) {
//       console.log("close clicked");
//       let modal = document.getElementById("myModal");
//       modal.style.display = "none";
//     },
//     onOutsideClick(event) {
//       console.log("outside clicked");
//       let modal = document.getElementById("myModal");
//       if (event.target == modal) {
//         modal.style.display = "none";
//       }
//     },
//     })),
//     Alpine.data('hdrData', () => ({
//       btn: document.getElementById("myBtn"),
//       span: document.getElementsByClassName("close")[0],
//       // Functions
//       onClick(event) {
//         let modal = document.getElementById("myDropdown");
//         modal.style.display = "block";
//       },
//       onElementClick(event) {
//         // Disappear the dropdown
//         let modal = document.getElementsByClassName("hdr__dropdown-content")[0];
//         modal.style.display = "none";
//         // let sidnav = document.getElementsByClassName("sidenav")[0];
//         // sidnav.style.display = "block";
//       },
//       onOutsideClick(event) {
//         console.log("outside clicked");
//         let modal = document.getElementById("myDropdown");
//         modal.style.display = "none";
//       },
//     })),
//     Alpine.data('settings', () => ({
//       active: document.getElementsByClassName("page__general")[0],
//       parm2: "",
//       onMenuClick(event) {
//         let text = event.target.innerText;
//         console.log("menu clicked: ", text);

//         if (text === "General") {
//           console.log("General clicked");
//           document.getElementsByClassName("page__general")[0].style.display = "flex";
//           document.getElementsByClassName("page__test")[0].style.display = "none";
//           document.getElementsByClassName("page__unigydb")[0].style.display = "none";
//         }

//         if (text === "Test") {
//           console.log("Test clicked");      
//           document.getElementsByClassName("page__test")[0].style.display = "flex";
//           document.getElementsByClassName("page__general")[0].style.display = "none";
//           document.getElementsByClassName("page__unigydb")[0].style.display = "none";
          
//         }

//         if (text === "Unigy Database") {
//           console.log("Unigy Database clicked");
//           document.getElementsByClassName("page__unigydb")[0].style.display = "flex";
//           document.getElementsByClassName("page__general")[0].style.display = "none";
//           document.getElementsByClassName("page__test")[0].style.display = "none";
//         }
//       },
//     })),
//     Alpine.data('sidenav', () => ({
//       list_search: "",
//       chevronRotated: false,
//       nodeListCopy: [],
//       strLength: 0,

//       onElementClick(el) {
//         let children = el.childNodes;

//         children.forEach((element) => {
//           if (element.tagName === "I") {
//             element.className =
//               element.className === "fa fa-chevron-right rotate_back"
//                 ? "fa fa-chevron-right rotate_fwd"
//                 : "fa fa-chevron-right rotate_back";
//           }
//         });

//         this.nodeListCopy = [];
//       },
//       onSearchChange(event) {
//         let el = document.getElementsByClassName("list_container__ul")[0];
//         const list = el.childNodes;

//         // copy original list to nodeListCopy
//         if (this.nodeListCopy.length === 0) {         
//           list.forEach(function (currentValue, currentIndex, listObj) {           
//             let copy = currentValue.cloneNode(true);
//             this.nodeListCopy.push(copy);          
//           }, this);
//         }
        
//         // purge the list
//         while (el.firstChild) {
//           el.removeChild(el.firstChild);
//         }

//         // filter the list
//         this.nodeListCopy.forEach(function (currentValue, currentIndex, listObj) {
//           if (currentValue.tagName === "LI") {                       
//             if (currentValue.innerText.includes(this.list_search)) {              
//               el.appendChild(currentValue);
//             }
//           } else {
//             el.appendChild(currentValue);
//           }
//         }, this);        
//       },
//     })),
//     Alpine.data('lstable', () => ({
//       someVar: "",
//       info: 
//         { 
//           name: "", 
//           vip: "", 
//           ip: "", 
//           enterprise: "",
//           role: "", 
//           action: "" 
//         },
//         onHdrClick(event) {
//           let modal = document.getElementsByClassName("tbl-hdr-modal")[0];
//           modal.style.display = "flex";
//           modal.style.left = event.clientX - 250 + "px";
//           modal.style.top = event.clientY - 100 + "px";
//           let modalText = document.getElementsByClassName(
//             "tbl-hdr-modal__text"
//           )[0];
//           modalText.innerText = event.target.innerText;
//         },
//         onCloseClick(event) {
//           console.log("close clicked");
//           let modal = document.getElementsByClassName("tbl-hdr-modal")[0];
//           modal.style.display = "none";
//         },  
//         async onRowClick(el) {
//           console.log("row clicked: ", el);
//           const formData = new FormData(); 
//           const myHeaders = new Headers();       

//           let infoBox = document.getElementsByClassName("table-row-summary")[0];
//           infoBox.classList.add("table-row-summary__on");
        
//           let children = el.target.parentNode.childNodes;
//           children.forEach((element, i) => {
//             console.log("element: ", element.innerText + " - " + i);
//             switch (i) {
//               case 7:
//                 this.info.name = element.innerText;
//                 break;
//               case 11:
//                 this.info.ip = element.innerText;
//                 break;
//               case 13:
//                 this.info.vip = element.innerText;
//                 break;
//               case 21:
//                 this.info.enterprise = element.innerText;
//                 break;
//               case 23:
//                   this.info.role = element.innerText;
//                   break;            
//               case 25:
//                 this.info.action = element.innerText;
//                 break;
//             }          
//         })

//         myHeaders.append("Content-Type", "application/x-www-form-urlencoded");        
//         formData.append("view_id", "unigystatus");
//         formData.append("type", "request");
//         formData.append("target", "ip");
//         formData.append("data", this.info.ip);
        
//         const myRequest = new Request("/request/status", {
//           method: "POST",
//           headers: myHeaders,
//           body: new URLSearchParams(formData)
//         });

//         const response = await fetch(myRequest);
//         if (!response.ok) {
//           throw new Error(`Response status: ${response.status}`);
//         }

//         // console.log("response: ", response);

//         const json = await response.json();
//         console.log(json);
//         if (json.Server === "RUNNING") {
//           document.getElementById("info_box1").style.color = "green";
//         } else {
//           document.getElementById("info_box1").style.color = "red";
//         }
        
        

//       },

//       onUMSClick() {
//         console.log("UMS clicked");
//         var strWindowFeatures = "location=yes,height=570,width=520,scrollbars=yes,status=yes";
//         // var URL = "https://www.linkedin.com/cws/share?mini=true&amp;url=" + location.href;
//         let URL = `https://${this.info.vip}/ums2/index.html?UMSClient=`
//         var win = window.open(URL, "_blank");
        
        
//       },

//       onDetailsClick() {
//         console.log("Details clicked");
//         var strWindowFeatures = "location=yes,height=570,width=520,scrollbars=yes,status=yes";
//         // var URL = "https://www.linkedin.com/cws/share?mini=true&amp;url=" + location.href;
//         let URL = `https://${this.info.vip}/haservices/checkHAStatus`
//         var win = window.open(URL, "_blank");
//       },

//       onInfoLSideClick(ev) {
//         console.log("info clicked" + ev.target.innerText);
//         // navigator.clipboard.writeText(ev.target.innerText);
//         // alert("Copied the text: " + ev.target.innerText);
//         Window.navigator.clipboard.writeText(copyText).then(function() {
//           alert("Text copied to clipboard: " + copyText);
//         }).catch(function(error) {
//           alert("Failed to copy text: " + error);
//         });
        
//       },          
//     }))
// });


