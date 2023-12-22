import { useQuery } from "@tanstack/react-query";
import axios from "axios";

// @ts-ignore
const useIssues = () => {
	return useQuery({
		queryKey: ["api"],
		queryFn: () => axios.get<{ message: string }>("/api/issue"),
	});
};

export { useIssues };
