interface Invocation {
  id: string;
  client_name: string;
  function_name: string;
  image: string;
  started_time: string;
  ended_time: string;
  input: string;
  output: string;
}

export default Invocation;
