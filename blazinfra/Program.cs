using Microsoft.AspNetCore.Components;
using Microsoft.AspNetCore.Components.Web;
using BlazInfra.Data;
using BlazInfra.Protobuf;
using Grpc.Net.Client;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddRazorPages();
builder.Services.AddServerSideBlazor();
builder.Services.AddSingleton<WeatherForecastService>();
builder.Services.AddGrpc();
builder.Services.AddBlazorBootstrap();

builder.Services.AddSingleton<MonitoringService.MonitoringServiceClient>(services =>
{
    var channel = GrpcChannel.ForAddress("http://localhost:50051");
    return new MonitoringService.MonitoringServiceClient(channel);
}
);

var app = builder.Build();

// Configure the HTTP request pipeline.
if (!app.Environment.IsDevelopment())
{
    app.UseExceptionHandler("/Error");
    // The default HSTS value is 30 days. You may want to change this for production scenarios, see https://aka.ms/aspnetcore-hsts.
    app.UseHsts();
}

app.UseHttpsRedirection();

app.UseStaticFiles();

app.UseRouting();

app.MapBlazorHub();
app.MapFallbackToPage("/_Host");

app.Run();
