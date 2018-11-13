
	function SendRequest(url_target,pesan,onResponse) {
		var xhttp = new XMLHttpRequest();
		xhttp.onreadystatechange = function() {
    			if (this.readyState == 4 && this.status == 200) {
					onResponse(JSON.parse(this.responseText));
				}
  		};

		xhttp.open("POST", url_target, true);
  		xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
  		xhttp.send(pesan);
	}

	function UploadFile(formName,files,url,onProgress,onResponse){

			var file = files[0];
			var fd = new FormData();
			fd.append(formName, file);
	
			var xhr = new XMLHttpRequest();
			xhr.open('POST', url, true);
	
			xhr.upload.onprogress = function(e) {
				  if (e.lengthComputable) {
						var percentComplete = (e.loaded / e.total) * 100;
						onProgress(percentComplete + '% uploaded');
				  }
			};
			xhr.onload = function() {
				  if (this.status == 200) {
						var resp = JSON.parse(this.response);
						onResponse(resp)
				  };
			};
			xhr.send(fd);
	}




	

     