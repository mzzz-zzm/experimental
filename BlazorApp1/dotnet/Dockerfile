FROM mcr.microsoft.com/dotnet/sdk:7.0 as build
WORKDIR /app

COPY BlazorApp1.sln ./
COPY BlazorApp1/BlazorApp1.csproj ./BlazorApp1/

RUN dotnet restore
COPY . ./
RUN dotnet publish -c Release -o out

FROM nginx:1.23.0-alpine
WORKDIR /app
EXPOSE 8080
ARG ENVIRONMENT
COPY nginx.conf /etc/nginx/nginx.conf
RUN sed -i "s/envname/${ENVIRONMENT}/" /etc/nginx/nginx.conf
COPY --from=build /app/out/wwwroot /usr/share/nginx/html
