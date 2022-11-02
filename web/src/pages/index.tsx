// interface HomeProps {
//   count: number;
// }

// export default function Home(props: HomeProps) {
//   return (
//     <h1 className="text-violet-500 font-bold text-4xl">
//       Contagem: {props.count}
//     </h1>
//   );
// }

// export const getServerSideProps = async () => {
//   const response = await fetch("http://localhost:8080/pools/count");
//   const data = await response.json();

//   return {
//     props: {
//       count: data.count,
//     },
//   };
// };

import appPreviewImg from "../assets/app-nlw-copa-preview.png";
import logoImg from "../assets/logo.svg";
import iconCheckImg from "../assets/icon-check.svg";

import Image from "next/image";

export default function Home() {
  return (
    <div>
      <main>
        <Image src={logoImg} alt="NLW Logo"></Image>

        <h1>Crie seu bolão da copa</h1>

        <form>
          <input type="text" required placeholder="Qual o nome do seu bolão?" />

          <button type="submit">Criar meu bolão</button>
          <div>
            <div>
              <Image src={iconCheckImg} alt=""></Image>
            </div>
            <div></div>
          </div>
        </form>
      </main>
      <Image
        src={appPreviewImg}
        alt="Dois celulares exibindo uma previa da aplicação do bolão"
        quality={100}
      />
    </div>
  );
}
