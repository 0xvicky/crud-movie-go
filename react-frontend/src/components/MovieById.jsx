import React, {useState} from "react";

const MovieById = ({setData}) => {
  const [id, setId] = useState();

  const getMovie = async () => {
    try {
      const res = await fetch("http://192.168.1.11:8080/movies/1");
      const data = await res?.json();
      console.log(data);
      //   console.log(get);
      setData(data);
    } catch (error) {
      console.log(`Error occured:${error}`);
    }
  };

  return (
    <div className='flex'>
      <input
        type='text'
        placeholder='Id'
        onChange={e => setId(e.target.value)}
      />
      <button
        className='bg-black text-purple-100 p-1 rounded-md'
        onClick={getMovie}>
        Get Movie By id
      </button>
    </div>
  );
};

export default MovieById;
