
import { useState, useRef } from "react";

const Blog = () => {
  const [inputs, setInputs] = useState([]); // Array of sequential inputs (text, images, videos)
  const fileInputRef = useRef([]); // Ref to store files temporarily

  // Function to handle text input changes
  const handleTextChange = (index, value) => {
    const updatedInputs = inputs.map((input, i) =>
      i === index ? { ...input, value } : input
    );
    setInputs(updatedInputs); // Update the input array with new text value
  };

  // Function to add new text input
  const addTextInput = () => {
    setInputs([...inputs, { type: 'text', value: '' }]); // Add a new text input to the array
  };

  // Function to handle file changes and store them in sequence
  const handleFileChange = (e) => {
    const files = Array.from(e.target.files); // Convert FileList to Array

    files.forEach((file) => {
      const fileUrl = URL.createObjectURL(file); // Generate a URL for display
      const fileType = file.type.startsWith('image') ? 'image' : 'video'; // Determine if it's image or video
      setInputs((prevInputs) => [...prevInputs, { type: fileType, value: fileUrl, file }]); // Append file in sequence
    });
  };

  // Function to delete an input (text, image, or video) based on its index
  const handleDelete = (index) => {
    const updatedInputs = inputs.filter((_, i) => i !== index); // Filter out the item at the given index
    setInputs(updatedInputs); // Update the input array after deletion
  };

  // Function to handle form submission (sending files to the backend)
  const handleSubmit = async () => {
    const formData = new FormData();
    
    // Append each item from the input array (text, images, videos) to formData
    inputs.forEach((input) => {
      if (input.type === 'text') {
        formData.append('text', input.value); // Append text
      } else {
        formData.append('files', input.file); // Append file (image or video)
      }
    });

    // Example: sending formData to the backend
    try {
      const response = await fetch('/api/upload', {
        method: 'POST',
        body: formData,
      });
      if (response.ok) {
        alert('Files uploaded successfully!');
      } else {
        alert('Upload failed.');
      }
    } catch (error) {
      console.error('Error uploading files:', error);
    }
  };

  return (
    <div className="p-4">
      {/* Navigation bar with buttons to add text and upload files */}
      <nav className="flex gap-4 mb-4">
        <button
          onClick={addTextInput}
          className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600"
        >
          Add Text
        </button>
        <button className="bg-green-500 text-white px-4 py-2 rounded-md hover:bg-green-600">
          <label htmlFor="file-upload" className="cursor-pointer">
            Upload Image/Video
          </label>
        </button>
        <input
          id="file-upload"
          type="file"
          style={{ display: 'none' }}
          accept="image/*,video/*"
          multiple // Allow multiple files
          onChange={handleFileChange} // Handle file selection
        />
      </nav>

      {/* Display inputs in sequence */}
      <div>
        {inputs.map((input, index) => (
          <div key={index} className="border border-gray-300 p-4 my-4 rounded-lg">
            {input.type === 'text' && (
              <textarea
                value={input.value}
                onChange={(e) => handleTextChange(index, e.target.value)}
                placeholder="Enter text here..."
                rows={3}
                className="w-full p-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            )}
            {input.type === 'image' && (
              <div>
                <img src={input.value} alt={`Uploaded ${input.type}`} className="max-w-full h-auto" />
                <button
                  onClick={() => handleDelete(index)}
                  className="bg-red-500 text-white px-2 py-1 mt-2 rounded-md hover:bg-red-600"
                >
                  Delete Image
                </button>
              </div>
            )}
            {input.type === 'video' && (
              <div>
                <video controls className="max-w-full h-auto">
                  <source src={input.value} type="video/mp4" />
                  Your browser does not support the video tag.
                </video>
                <button
                  onClick={() => handleDelete(index)}
                  className="bg-red-500 text-white px-2 py-1 mt-2 rounded-md hover:bg-red-600"
                >
                  Delete Video
                </button>
              </div>
            )}
          </div>
        ))}
      </div>

      {/* Submit button to send data to the backend */}
      <button
        onClick={handleSubmit}
        className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600"
      >
        Submit
      </button>
    </div>
  );
};

export default Blog;


